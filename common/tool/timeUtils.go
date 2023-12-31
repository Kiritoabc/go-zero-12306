package tool

import (
	"go-zero-12306/common/constant"
	"time"
)

// TomorrowDateStr 获取明天的日期
func TomorrowDateStr() string {
	// 获取当前时间
	now := time.Now()
	// 计算，明天时间
	tomorrow := now.AddDate(0, 0, 1)
	return tomorrow.Format(constant.Datetimeform1)
}
