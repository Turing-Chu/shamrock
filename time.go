// Author: Turing Zhu
// Date: 6/15/21 9:53 AM
// File: time.go

package shamrock 

import (
	"time"
)

func FormatStdTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func FormatTimeWitNano(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.999999")
}
