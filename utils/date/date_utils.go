package date

import "time"

const (
	apiDateFormat = "2006-01-02T15:04:05Z"
)

func GetNowString() string {
	return GetNow().Format(apiDateFormat)
}

func GetNow() time.Time {
	return time.Now().UTC()
}
