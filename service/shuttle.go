package service

import (
	"awesomeProject/api"
	"fmt"
	"sync"
	"time"

	"awesomeProject/config"
	"awesomeProject/models"
	"awesomeProject/utils"
)

// DisplayShuttleList 显示班车列表
func DisplayShuttleList(shuttleList *models.ShuttleList) {
	fmt.Println("班车列表：")
	fmt.Println("序号 | 类型     | 发车时间")
	fmt.Println("-----|----------|----------")
	for i, shuttle := range shuttleList.Data {
		fmt.Printf("[%d]  | %s | %s\n",
			i,
			utils.GetShuttleType(shuttle.Type),
			shuttle.OriginTime)
	}
}

// DisplayReservedSeats 显示座位预定状态
func DisplayReservedSeats(shuttle models.ShuttleRoute, reservedSeats *models.ReservedSeats) {
	fmt.Println("\n座位预定状态：")
	fmt.Println("----------------------------------------")
	fmt.Printf("类型：%s\n", utils.GetShuttleType(shuttle.Type))
	fmt.Printf("时间：%s - %s\n", shuttle.OriginTime, shuttle.EndTime)
	fmt.Printf("起点：%s\n", shuttle.OriginAddress)
	fmt.Printf("终点：%s\n", shuttle.EndAddress)
	fmt.Printf("已预订座位：%v\n", reservedSeats.Data.ReservedSeatNumber)
	availableSeats := utils.GetAvailableSeats(reservedSeats.Data.ReservedSeatNumber, config.AppConfig.TotalSeats)
	fmt.Printf("空余座位： %v\n", availableSeats)
	fmt.Printf("班车状态： %s\n", utils.IfFull(reservedSeats.Data.IsFull))
	fmt.Println("----------------------------------------")
}

// WaitForOrderTime 等待到达下单时间
func WaitForOrderTime(originTime time.Time) {
	waitDuration := utils.CalculateWaitTime(originTime)

	if waitDuration > 0 {
		fmt.Printf("班车发车时间: %s, 将在发车前一小时 (%s) 开始下单\n",
			originTime.Format("2006-01-02 15:04"),
			originTime.Add(-1*time.Hour).Format("2006-01-02 15:04"))

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		targetTime := originTime.Add(-1 * time.Hour)

		for range ticker.C {
			now := time.Now()
			remaining := targetTime.Sub(now)
			if remaining <= 0 {
				fmt.Println("到达下单时间，开始下单流程...")
				break
			}
			fmt.Printf("\r当前时间: %s | 距离下单时间还剩: %s ",
				now.Format("2006-01-02 15:04:05"),
				utils.FormatDuration(remaining))
		}
	} else {
		fmt.Println("当前时间已超过下单时间，直接开始下单流程...")
	}
}

// ProcessOrder 处理下单流程
func ProcessOrder(shuttle models.ShuttleRoute, date string) {
	mailService := NewMailService()

	for {
		reservedSeats, err := api.GetReservedSeats(shuttle.ID, date, config.AppConfig.UserID)
		if err != nil {
			fmt.Println("获取座位预定状态错误:", err)
			return
		}
		availableSeats := utils.GetAvailableSeats(reservedSeats.Data.ReservedSeatNumber, config.AppConfig.TotalSeats)
		fmt.Printf("更新后的空余座位: %v\n", availableSeats)

		if len(availableSeats) < 3 {
			fmt.Println("可用座位不足3个，只尝试可用的座位")
		}

		var wg sync.WaitGroup
		orderSuccess := false
		successfulSeat := 0
		emailSent := false
		var mutex sync.Mutex

		orderCount := 3
		if len(availableSeats) < orderCount {
			orderCount = len(availableSeats)
		}

		if orderCount == 0 {
			fmt.Println("当前没有可用座位，等待5秒后重试...")
			time.Sleep(5 * time.Second)
			continue
		}

		for i := 0; i < orderCount; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				message, err := api.CreateOrder(shuttle.ID, date, config.AppConfig.UserID, int64(availableSeats[i]))
				if err != nil {
					fmt.Println("下单失败:", err)
				} else {
					fmt.Printf("%d号座位下单成功: %s\n", availableSeats[i], *message)

					mutex.Lock()
					if !orderSuccess {
						orderSuccess = true
						successfulSeat = availableSeats[i]

						// 只发送一次邮件通知
						if !emailSent {
							userEmail := getUserEmail()
							if userEmail != "" {
								err := mailService.SendBookingSuccessEmail(shuttle, date, successfulSeat, userEmail)
								if err != nil {
									fmt.Printf("邮件发送失败: %v\n", err)
								} else {
									fmt.Printf("购票成功邮件已发送至: %s\n", userEmail)
									emailSent = true
								}
							}
						}
					}
					mutex.Unlock()
				}
			}(i)
		}
		wg.Wait()
		if orderSuccess {
			break
		}
		fmt.Println("本轮下单失败，继续尝试...")
		time.Sleep(2 * time.Second)
	}
}

// getUserEmail 获取用户邮箱地址
func getUserEmail() string {
	return config.AppConfig.Mail.UserEmail
}
