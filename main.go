package main

import (
	"flag"
	"fmt"
	"time"

	"awesomeProject/api"
	"awesomeProject/config"
	"awesomeProject/service"
	"awesomeProject/utils"
)

func main() {
	// 初始化配置
	config.LoadConfig()

	indexPtr := flag.Int("index", 0, "指定要购买的班车序号（从0开始）")
	addressPtr := flag.Int("address", 0, "指定校车路线（0：中关村->良乡，1：良乡->中关村）")
	dateAddrPtr := flag.Int("date", 0, "指定日期为今日后的第几天（默认0为今日）")
	flag.Parse()

	address := "中关村校区->良乡校区"
	if *addressPtr == 1 {
		address = "良乡校区->中关村校区"
	} else if *addressPtr != 0 {
		fmt.Printf("错误：无效的路线参数 %d，可选值为 0 或 1\n", *addressPtr)
		return
	}

	defaultDate := time.Now().AddDate(0, 0, *dateAddrPtr).Format("2006-01-02")
	shuttleList, err := api.GetShuttleList(defaultDate, address)
	if err != nil {
		fmt.Println("获取班车列表错误:", err)
		return
	}

	service.DisplayShuttleList(shuttleList)

	if *indexPtr < 0 || *indexPtr >= len(shuttleList.Data) {
		fmt.Printf("错误：无效的班车序号 %d，可选范围为 0-%d\n", *indexPtr, len(shuttleList.Data)-1)
		return
	}

	if len(shuttleList.Data) > 0 {
		shuttle := shuttleList.Data[*indexPtr]
		if shuttle.Type != 0 {
			fmt.Println("只能抢普通校车，不能抢彩虹巴士...")
			return
		}

		// 显示选中班车的详细信息
		fmt.Printf("\n已选择班车 [%d]：\n", *indexPtr)
		reservedSeats, err := api.GetReservedSeats(shuttle.ID, defaultDate, config.AppConfig.UserID)
		if err != nil {
			fmt.Println("获取座位预定状态错误:", err)
			return
		}
		service.DisplayReservedSeats(shuttle, reservedSeats)

		originTime, err := utils.ParseTime(defaultDate, shuttle.OriginTime)
		if err != nil {
			fmt.Println("解析发车时间错误:", err)
			return
		}

		service.WaitForOrderTime(originTime)

		service.ProcessOrder(shuttle, defaultDate)
	}
}
