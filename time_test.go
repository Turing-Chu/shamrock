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

func TestFormatTimeWitMicro(t *testing.T) {
	now := time.Now()
	fmt.Printf("now: %s\tForamtStdTimeWithMicro: %s\t now.Format: %s\n", now, FormatTimeWitMicro(now), now.Format("2006-01-02 15:04:05.999999"))
	if FormatTimeWitMicro(now) != now.Format("2006-01-02 15:04:05.999999") {
		t.Fatalf("now:%s, ForamtStdTimeWithMicro:%s, now.Format: %s\n",
			now, FormatTimeWitMicro(now), now.Format("2006-01-02 15:04:05.999999"))
	}
}

func TestFormatTimeWitNano(t *testing.T) {
	now := time.Now()
	fmt.Printf("now: %s\tForamtStdTimeWithNano: %s\t now.Format: %s\n", now, FormatTimeWitNano(now), now.Format("2006-01-02 15:04:05.999999999"))
	if FormatTimeWitNano(now) != now.Format("2006-01-02 15:04:05.999999999") {
		t.Fatalf("now:%s, ForamtStdTimeWithNano:%s, now.Format: %s\n",
			now, FormatStdTime(now), now.Format("2006-01-02 15:04:05.999999999"))
	}
}

func TestFormatDuration(t *testing.T) {
	testData := []time.Duration{
		0,
		time.Second,
		time.Minute,
		time.Hour,
		time.Hour * 24,
		time.Hour * 23,
		time.Hour * 25,
		time.Hour * 34,
		time.Hour * 123,
		time.Hour * 12300,
		-time.Hour,
	}
	for _, v := range testData {
		fmt.Printf("duration: %s,\tformat: %s\n", v, FormatDuration(v))
	}
}
