package services

import "time"

func GetCurrentTime() string {
	return time.Now().String()
}
