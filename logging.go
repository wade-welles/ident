package main

import (
	"fmt"
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

func logInfo(format string, args ...interface{}) {
	logMessage("INFO", format, args...)
}

func logError(format string, args ...interface{}) {
	logMessage("ERROR", format, args...)
}

func logMessage(level, format string, args ...interface{}) {
	msg := fmt.Sprintf("[%s] %s", level, format)
	Logger.Output(3, fmt.Sprintf(msg, args...))
}
