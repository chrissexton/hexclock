// A simple library to convert time objects into string representations of hexadecimal time.
// See: https://en.wikipedia.org/wiki/Florence_Mean_Time
package hextime

import (
	"time"
	"fmt"
)

// Return current hex time
func Now() string {
	return Convert(time.Now())
}

// Convert a given time into a hex time
func Convert(now time.Time) string {
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).UnixNano()
	current := now.UnixNano()

	sinceMidnight := current - today

	hexSecond := int64(1.318359375 * 1000000000)

	hexSeconds := sinceMidnight / hexSecond

	hexHour := hexSeconds / 16 / 16 / 16 % 16
	hexMaxime := hexSeconds / 16 / 16 % 16
	hexMinute := hexSeconds / 16 % 16
	hexSeconds = hexSeconds % 16

	return fmt.Sprintf(",%x%x%x%x", hexHour, hexMaxime, hexMinute, hexSeconds)
}
