package z

import "time"

const (
	DateTime = "2006-01-02 15:04:05"
	Date     = "2006-01-02"
	Time     = "15:04:05"
)

func Now() {
	time.Now().Format(DateTime)
}
