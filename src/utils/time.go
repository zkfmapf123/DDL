package utils

import "time"

// YYYY-MM-DD
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02")
}
