package hextime

import (
	"time"
	"fmt"
)

func getTime(when string) time.Time {
	t, _ := time.Parse(time.Kitchen, when)
	return t
}

func Now() string {
	return Convert(time.Now())
}

func Convert(now time.Time) string {
	now = now.UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).UnixNano()
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
