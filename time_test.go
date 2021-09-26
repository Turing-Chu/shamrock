// Author: Turing Zhu
// Date: 7/23/21 2:33 PM
// File: time_test.go

package shamrock

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatStdTime(t *testing.T) {
	now := time.Now()
	fmt.Printf("now: %s\tFormatStdTime: %s\t now.Format: %s\n", now, FormatStdTime(now), now.Format("2006-01-02 15:04:05"))
	if FormatStdTime(now) != now.Format("2006-01-02 15:04:05") {
		t.Fatalf("now:%s, ForamtStdTime:%s, now.Format: %s\n",
			now, FormatStdTime(now), now.Format("2006-01-02 15:04:05"))
	}
}

func TestFormatTimeWithMilli(t *testing.T) {
	now := time.Now()
	fmt.Printf("now: %s\tForamtStdTimeWithMilli: %s\t now.Format: %s\n", now, FormatTimeWithMilli(now), now.Format("2006-01-02 15:04:05.999"))
	if FormatTimeWithMilli(now) != now.Format("2006-01-02 15:04:05.999") {
		t.Fatalf("now:%s, ForamtStdTimeWithMilli:%s, now.Format: %s\n",
			now, FormatTimeWithMilli(now), now.Format("2006-01-02 15:04:05.999"))
	}
}

func TestFormatTimeWithMicro(t *testing.T) {
	now := time.Now()
	fmt.Printf("now: %s\tForamtStdTimeWithMicro: %s\t now.Format: %s\n", now, FormatTimeWithMicro(now), now.Format("2006-01-02 15:04:05.999999"))
	if FormatTimeWithMicro(now) != now.Format("2006-01-02 15:04:05.999999") {
		t.Fatalf("now:%s, ForamtStdTimeWithMicro:%s, now.Format: %s\n",
			now, FormatTimeWithMicro(now), now.Format("2006-01-02 15:04:05.999999"))
	}
}

func TestFormatTimeWithNano(t *testing.T) {
	now := time.Now()
	fmt.Printf("now: %s\tForamtStdTimeWithNano: %s\t now.Format: %s\n", now, FormatTimeWithNano(now), now.Format("2006-01-02 15:04:05.999999999"))
	if FormatTimeWithNano(now) != now.Format("2006-01-02 15:04:05.999999999") {
		t.Fatalf("now:%s, ForamtStdTimeWithNano:%s, now.Format: %s\n",
			now, FormatStdTime(now), now.Format("2006-01-02 15:04:05.999999999"))
	}
}

// 测试文件
func TestFormatDuration(t *testing.T) {
	fmt.Println("TEST FormatDuration")
	durations := []time.Duration{
		0,
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
