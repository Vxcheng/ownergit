package main

import (
	"fmt"
)

func main() {
	overtime(August)
}

// 每天打卡时间需满12个小时， 周五减一个小时, 精确到分钟
type MonthDays int

const (
	August MonthDays = 21
)

func overtime(m MonthDays) {
	switch m {
	case August:
		m.August()
	}
}

func (m MonthDays) August() {
	needTime := int(m)*12*60 - 4*60 // 应出勤
	offTime := 1 * 8 * 60           // 请假时间
	needTime -= offTime
	totalTime := 0 // 实际出勤
	totalTime += (12*60 + 40) + (12*60 + 54) + (12*60 + 77) + (12*60 + 16) + (10*60 + 11) +
		+(12*60 + 121) + (12*60 + 20) + (12*60 + 46) + (12*60 + 37) + (12*60 + 17) +
		+(12*60 + 29) + (12*60 + 65) + (12*60 + 44) + (12*60 + 64) +
		+(12*60 + 94) + (12*60 + 22) + (12*60 + 77) + (12*60 + 17) + (11*60 + 38) +
		+(9*60 + 12)
	sub := totalTime - needTime
	fmt.Printf("totalTime-needTime = %d, about to '%.2fh'.\n", sub, float64(sub/60))
	if sub > 0 {
		fmt.Println("ok!")
	}
}
