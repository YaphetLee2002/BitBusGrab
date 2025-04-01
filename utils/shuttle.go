package utils

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

// GetAvailableSeats 获取空余座位
func GetAvailableSeats(reservedSeatNumber []string, totalSeats int) []int {
	reservedMap := make(map[int]bool)
	for _, seatStr := range reservedSeatNumber {
		seatNum, err := strconv.Atoi(seatStr)
		if err == nil {
			reservedMap[seatNum] = true
		}
	}

	var availableSeats []int
	for i := 3; i <= totalSeats; i++ {
		if i == 49 {
			continue
		}
		if !reservedMap[i] {
			availableSeats = append(availableSeats, i)
		}
	}

	sort.Ints(availableSeats)

	return availableSeats
}

// ParseTime 解析时间字符串为time.Time
func ParseTime(dateStr, timeStr string) (time.Time, error) {
	dateTimeStr := fmt.Sprintf("%s %s", dateStr, timeStr)
	return time.ParseInLocation("2006-01-02 15:04", dateTimeStr, time.Local)
}

// CalculateWaitTime 计算等待时间
func CalculateWaitTime(originTime time.Time) time.Duration {
	orderTime := originTime.Add(-1 * time.Hour)
	now := time.Now()
	fmt.Println("当前时间:", now.Format("2006-01-02 15:04:05"))
	if now.After(orderTime) {
		return 0
	}

	waitTime := orderTime.Sub(now)
	if waitTime <= 5*time.Second {
		return 0
	}
	return waitTime - 5*time.Second
}

// GetShuttleType 获取班车类型名称
func GetShuttleType(shuttleType int) string {
	switch shuttleType {
	case 0:
		return "校车"
	case 1:
		return "彩虹巴士"
	default:
		return "未知类型"
	}
}

// IfFull 判断班车是否已满
func IfFull(isFull int) string {
	if isFull == 1 {
		return "已满"
	}
	return "未满"
}

// FormatDuration 格式化持续时间
func FormatDuration(d time.Duration) string {
	h := d / time.Hour
	d = d % time.Hour
	m := d / time.Minute
	d = d % time.Minute
	s := d / time.Second

	if h > 0 {
		return fmt.Sprintf("%d小时%d分钟%d秒", h, m, s)
	} else if m > 0 {
		return fmt.Sprintf("%d分钟%d秒", m, s)
	}
	return fmt.Sprintf("%d秒", s)
}
