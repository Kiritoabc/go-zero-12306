package tool

import (
	"go-zero-12306/common/constant"
	"time"
)

// TomorrowDateStr 获取明天的日期
func TomorrowDateStr() string {
	tomorrow := TomorrowDate()
	return tomorrow.Format(constant.Datetimeform1)
}

func TomorrowDate() time.Time {
	// 获取当前时间
	now := time.Now()
	// 计算，明天时间
	return now.AddDate(0, 0, 1)
}

func ConvertDateToLocalTime(dateStr string, format string) string {
	// 在这里，假设 dateStr 是符合 "2006-01-02 15:04" 格式的字符串，需要使用对应的日期解析库进行转换（如 time 或者第三方库）
	departureTime, _ := time.Parse("2006-01-02 15:04", dateStr)
	return departureTime.Format(format)
}

func CalculateHourDifference(departureTimeStr, arrivalTimeStr string) int {
	departureTime, _ := time.Parse("2006-01-02 15:04", departureTimeStr)
	arrivalTime, _ := time.Parse("2006-01-02 15:04", arrivalTimeStr)
	duration := int(arrivalTime.Sub(departureTime).Hours())
	return duration
}
