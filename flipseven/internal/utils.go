package internal

import (
	"log"
)

func CustomPanic(err error, message string) {
	if err != nil {
		log.Panicf("%v: %v", message, err)
	}
}
