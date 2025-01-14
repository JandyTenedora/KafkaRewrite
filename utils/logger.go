package utils

import (
	"fmt"
	"log"
	"time"
)

type LogLevel int

const (
	INFO LogLevel = iota
	ERROR
	DEBUG
)

func Log(level LogLevel, message string) {
	timestamp := time.Now().Format("2000-01-01 00:00:00")

	var prefix string
	switch level {
	case INFO:
		prefix = "[INFO]"
	case ERROR:
		prefix = "[ERROR]"
	case DEBUG:
		prefix = "[DEBUG]"
	default:
		prefix = "[UNKNOWN]"
	}

	log.Println(fmt.Sprintf("%s: %s - %s", prefix, message, timestamp))
}
