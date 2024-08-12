package utils

import "time"

func GetCurrentTime() time.Time {
	return time.Now()
}

func GetCurrentTs() int64 {
	return GetCurrentTime().Unix()
}

func GetCurrentMs() int64 {
	return GetCurrentTime().UnixMilli()
}
