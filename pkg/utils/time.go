package utils

import (
	"fmt"
	"time"
)

func GetDurationFormatString(duration time.Duration) string {
	// hours
	hours := duration / time.Hour
	duration -= hours * time.Hour
	// minutes
	minutes := duration / time.Minute
	duration -= minutes * time.Minute
	// seconds
	seconds := duration / time.Second
	duration -= seconds * time.Second
	// milliseconds
	milis := duration / time.Millisecond

	return fmt.Sprintf("%02d:%02d:%02d.%03d", hours, minutes, seconds, milis)
}
