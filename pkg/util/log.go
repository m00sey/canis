package util

import (
	"log"
	"time"
)

func Logger(err error, _ time.Duration) {
	log.Println(err)
}
