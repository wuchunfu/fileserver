package utils

import (
	"strconv"
	"time"
)

// 格式化当前时间
func FormatNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 格式化时间
func FormatDateTime(datetime time.Time) string {
	return datetime.Format("2006-01-02 15:04:05")
}

// 时间转字符串 参数 true 精确到时分秒 false 精确到日期
func TimeToString(accurate bool) string {
	var timeLayout string
	if accurate == true {
		// 时间模板-精确
		timeLayout = "2006-01-02 15:04:05"
	} else {
		timeLayout = "2006-01-02"
	}
	// 当前时间
	nowTime := time.Now().Unix()
	// 转换当前时间戳为时间模板格式
	dateTime := time.Unix(nowTime, 0).Format(timeLayout)
	// 返回时间字符串
	return dateTime
}

// 将字符串类型的时间戳参数转为时间字符串
func UnixTimeToString(stamp string) string {
	stamp = stamp[:10]
	base, _ := strconv.ParseInt(stamp, 10, 64)
	timeLayout := "2006-01-02 15:04:05"
	dateTime := time.Unix(base, 0).Format(timeLayout)
	return dateTime
}

// 获取当前年份 月份 日期
func EnumerateDate() (year, month, day string) {
	year = strconv.Itoa(time.Now().Year())
	month = strconv.Itoa(int(time.Now().Month()))
	day = strconv.Itoa(time.Now().Day())
	return
}
