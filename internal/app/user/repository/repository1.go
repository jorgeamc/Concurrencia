package repository

import (
	"time"
)

func GetRepositoryOne() string {
	time.Sleep(5 * time.Second)
	return "Response Repository 1"
}
