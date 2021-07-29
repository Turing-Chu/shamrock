// Author: Turing Zhu
// Date: 6/15/21 9:53 AM
// File: time.go

package shamrock 

import (
	"fmt"
	"time"
)

const (
	StdTimeFmt      = "2006-01-02 15:04:05"
	StdMicroTimeFmt = "2006-01-02 15:04:05.999999"
)

func FormatStdTime(t time.Time) string {
	return t.Format(StdTimeFmt)
}

func FormatTimeWitNano(t time.Time) string {
	return t.Format(StdMicroTimeFmt)
}

func FormatDuration(d time.Duration) (result string) {
	days := d.Nanoseconds() / (24 * time.Hour.Nanoseconds())
	d = time.Duration(int64(d) % int64(24*time.Hour))
	hours := d / time.Hour
	d = d % time.Hour
	minutes := d / time.Minute
	d = d % time.Minute
	seconds := d / time.Second
	if days > 0 {
		result = fmt.Sprintf("%dd", days)
	}
	if minutes > 0 && seconds > 0 || hours > 0 {
		result = fmt.Sprintf("%s%dh", result, hours)
	}
	if seconds > 0 || minutes > 0 {
		result = fmt.Sprintf("%s%dm", result, minutes)
	}
	if seconds >= 0 {
		result = fmt.Sprintf("%s%ds", result, seconds)
	}
	return result
}

