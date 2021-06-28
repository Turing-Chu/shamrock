// Author: Turing Zhu
// Date: 2021/06/28 14:31 PM
// File: time_test.go

package shamrock

import (
	"fmt"
	"testing"
    "time"
)

// 测试文件
func TestFormatDuration(t *testing.T) {
	fmt.Println("TEST FormatDuration")
	durations := []time.Duration{
		time.Second,
		3 * time.Second,
		time.Minute,
		65 * time.Second,
		59 * time.Minute,
		60 * time.Minute,
		24 * time.Hour,
		25 * time.Hour,
		43 * time.Hour,
		125 * 24 * time.Hour,
	}
	for _, d := range durations {
		fmt.Printf("test FormatDuration: %s\n", FormatDuration(d))
	}
	fmt.Println()
}

func TestFormatTime(t *testing.T) {
	fmt.Println("TEST FormatTime")
	gmt, _ := time.LoadLocation("GMT")
	testData := []time.Time{
		time.Now(),
		time.Date(1, 1, 1, 0, 0, 0, 0, time.Local),
		time.Unix(0, 0),
		time.Date(2006, 1, 2, 15, 4, 5, 0, gmt),
		time.Now().Add(time.Hour * 10000),
		time.Now().Add(-1000 * time.Hour),
		time.Now().Add(time.Hour * 100000),
	}
	for _, t := range testData {
		fmt.Printf("test FormatTime: %s\n", (t))
	}
	fmt.Println()
}

