package date_utils

import "time"

const (
	apiDateLayout = "2020-11-26T15:04:05Z")

func GetNow() time.Time {
	return time.Now()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
