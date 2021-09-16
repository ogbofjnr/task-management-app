package utils

import "time"

func PointerTime(time time.Time) *time.Time {
	return &time
}

func StrToPointer(str string) *string {
	return &str
}
