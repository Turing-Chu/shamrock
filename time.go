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
	StdMilliTimeFmt = "2006-01-02 15:04:05.999"
	StdMicroTimeFmt = "2006-01-02 15:04:05.999999"
	StdNanoTimeFmt  = "2006-01-02 15:04:05.999999999"
)

func FormatStdTime(t time.Time) string {
	return t.Format(StdTimeFmt)
}

func FormatTimeWithMilli(t time.Time) string {
	return t.Format(StdMilliTimeFmt)
}

func FormatTimeWithMicro(t time.Time) string {
	return t.Format(StdMicroTimeFmt)
}

func FormatTimeWithNano(t time.Time) string {
	return t.Format(StdNanoTimeFmt)
}

func FormatDuration(d time.Duration) (result string) {
	if d == 0 {
		return "0s"
	}
	positive := true
	if d < 0 {
		positive = false
		d = -d
	}
	days := d.Nanoseconds() / (24 * time.Hour.Nanoseconds())
	d = time.Duration(int64(d) % int64(24*time.Hour))
	if days > 0 {
		result = fmt.Sprintf("%dd", days)
		if d == 0 {
			if !positive {
				result = fmt.Sprintf("-%s", result)
			}
			return result
		}
	}

	hours := d / time.Hour
	d = d % time.Hour
	if hours > 0 {
		if result == "" {
			result = fmt.Sprintf("%s%dh", result, hours)
		} else {
			result = fmt.Sprintf("%s%02dh", result, hours)
		}
		if d == 0 {
			if !positive {
				result = "-" + result
			}
			return result
		}
	}

	minutes := d / time.Minute
	d = d % time.Minute
	if minutes > 0 {
		if result == "" {
			result = fmt.Sprintf("%s%dm", result, minutes)
		} else {
			result = fmt.Sprintf("%s%02dm", result, minutes)
		}
		if d == 0 {
			if !positive {
				result = "-" + result
			}
			return result
		}
	}

	seconds := d.String()
	if result != "" && 0 <= d && d < 10*time.Second {
		result = fmt.Sprintf("%s0%s", result, seconds)
	} else {
		result = fmt.Sprintf("%s%s", result, seconds)
	}
	if !positive {
		result = "-" + result
	}

	return result
}
