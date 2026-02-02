package repository

import (
	"time"
)

func GetRepositoryTwo() string {
	time.Sleep(2 * time.Second)
	return "Response Repository 2"
}
