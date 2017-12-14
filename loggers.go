package drawing

import (
	"log"
	"time"
)

// LogDuration ... Sample duration logger that uses the standard logger.
func LogDuration(operation string, taken time.Duration) {
	log.Printf("- %s - %s", operation, taken)
}

// NoDurationLog ... Sample duration logger that does nothing.
func NoDurationLog(operation string, taken time.Duration) {}
