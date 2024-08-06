package utils

import (
	"log"
)

// Info logs an informational message.
func Info(message string) {
	log.Printf("INFO: %s", message)
}
