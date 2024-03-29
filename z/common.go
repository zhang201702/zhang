package z

import "time"

const (
	DateTime = "2006-01-02 15:04:05"
	Date     = "2006-01-02"
	Time     = "15:04:05"
)

func Now() string {
	return time.Now().Local().Format(DateTime)
}

func GetTime(strDateTime string) time.Time {
	time, _ := time.ParseInLocation(DateTime, strDateTime, time.Local)
	return time
}
