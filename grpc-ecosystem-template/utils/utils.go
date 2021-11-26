package utils

import "time"

const TIME_FORMATER = "2006-01-02"

func FormatTimeToDate(t time.Time) string {
	return t.Format(TIME_FORMATER)
}

func ParseTime(str string) (time.Time, error) {
	return time.Parse(TIME_FORMATER, str)
}
