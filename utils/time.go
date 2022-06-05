package utils

import "time"

const Format24 = "2006-01-02 15:04:05"

func GetNowFMT() string {
	return time.Now().Format(Format24)
}
