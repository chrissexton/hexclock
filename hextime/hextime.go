// A simple library to convert time objects into string representations of hexadecimal time.
// See: https://en.wikipedia.org/wiki/Florence_Mean_Time
package hextime

import (
	"fmt"
	"strings"
	"time"
)

const (
	// Default format for times
	CommaNotation           = ",%a%b%c%d"
	PeriodNotation          = ".%A%B%C%D"
	BoardmanNotationSeconds = "%A_%B%C_%D"
	BoardmanNotation        = "%A_%B%C"
)

// Return current hex time with default format
func Now() string {
	return Convert(time.Now())
}

// Convert current time with a custom format string
func Nowf(format string) string {
	return Convertf(format, time.Now())
}

// Convert time with default format
func Convert(now time.Time) string {
	return Convertf(PeriodNotation, now)
}

// Convert a given time into a hex time
func Convertf(format string, now time.Time) string {
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).UnixNano()
	current := now.UnixNano()

	sinceMidnight := current - today

	hexSecond := int64(1.318359375 * 1000000000)

	hexSeconds := sinceMidnight / hexSecond

	hexHour := hexSeconds / 16 / 16 / 16 % 16
	hexMaxime := hexSeconds / 16 / 16 % 16
	hexMinute := hexSeconds / 16 % 16
	hexSeconds = hexSeconds % 16

	return hexFmt(format, hexHour, hexMaxime, hexMinute, hexSeconds)
}

// Replace format string entities with time parts
func hexFmt(format string, h, mx, m, s int64) string {
	output := format
	output = strings.Replace(output, "%A", fmt.Sprintf("%X", h), 1)
	output = strings.Replace(output, "%B", fmt.Sprintf("%X", mx), 1)
	output = strings.Replace(output, "%C", fmt.Sprintf("%X", m), 1)
	output = strings.Replace(output, "%D", fmt.Sprintf("%X", s), 1)

	output = strings.Replace(output, "%a", fmt.Sprintf("%x", h), 1)
	output = strings.Replace(output, "%b", fmt.Sprintf("%x", mx), 1)
	output = strings.Replace(output, "%c", fmt.Sprintf("%x", m), 1)
	output = strings.Replace(output, "%d", fmt.Sprintf("%x", s), 1)
	return output
}
