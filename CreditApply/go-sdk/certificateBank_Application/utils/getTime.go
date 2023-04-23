package utils

import (
	"time"
)

func GetTime() string {
	currentTime := time.Now()
	dateString := currentTime.Format("2006-01-02") // 使用时间模板将时间格式化为字符串
	return dateString
}
