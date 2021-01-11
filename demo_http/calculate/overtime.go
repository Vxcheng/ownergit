package main

import (
	"fmt"
)

func main() {
	overtime(November)
}

// 每天打卡时间需满12个小时， 周五减一个小时, 精确到分钟
type MonthDays int

const (
	August    MonthDays = 8
	September MonthDays = 9
	Octorbor MonthDays = 10
	November MonthDays = 11
)

func overtime(m MonthDays) {
	switch m {
	case August:
		m.August()
	case September:
		m.September()
	case Octorbor:
		m.Octorbor()
	case November:
		m.November()
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
	fmt.Printf("at 'August', totalTime-needTime = %d, about to '%.2fh'.\t", sub, float64(sub/60))
	if sub > 0 {
		fmt.Println(">>>>>>>>>>>> ok!")
	}
}

func (m MonthDays) September() {
	offTime := 1*3*60   // 请假时间
	totalTime := 0 // 实际出勤
	totalTime += (+38) + (+72) + (+26) + (-2*60 - 7) +
		+(1*60 + 17) + (+15) + (1*60 + 9) + (+25) + (-1*60 - 49)+
		+ (1*60+ 24) + (-45) + (1*60+26) + (1*60+29) + (-1*60-18) +
		+(10) + (2*60+5) + (1*60+56) + (1*60+4)+
		+(40)+ (44) + (28)+ (-5*60-30)
	sub := totalTime - offTime
	fmt.Printf("at 'September', aver time sub = %d, about to '%.2fh'.\t", sub, float64(sub/60))
	if sub > 0 {
		fmt.Println(">>>>>>>>>>>> ok!")
	}
}

func (m MonthDays) Octorbor() {
	offTime := 1*3*60   // 请假时间
	totalTime := 0 // 实际出勤
	totalTime += (+28) +
	 +(+34) + (1*60 + 48) + (1*60 + 20) + (25) + ( -2*60 + 9) +
	  +(1*60 - 6) + (+20) + (-2*60 +11)+
		+ (38) + (13) + (54) + (33)
	sub := totalTime - offTime
	fmt.Printf("at 'September', aver time sub = %d, about to '%.2fh'.\t", sub, float64(sub/60))
	if sub > 0 {
		fmt.Println(">>>>>>>>>>>> ok!")
	}
}

func  (m MonthDays) November() {
	offTime := 0   // 请假时间
	totalTime := 0 // 实际出勤
	totalTime += (+40)  + (+37-3*60) + (+32) + (+41- 3*60) + 
	 +(+60+28) + (1*60 + 45) + (+5) + (54) + ( -2*60 + 15) + (6*60+ 19)+
	  +(44) + (+18) + (1)+ (55) + (60+44)+
		+ (60+26) + (16) + (35) + (60+24)+ (-2*60+2) 
	sub := totalTime - offTime
	fmt.Printf("at 'September', aver time sub = %d, about to '%.2fh'.\t", sub, float64(sub/60))
	if sub > 0 {
		fmt.Println(">>>>>>>>>>>> ok!")
	}
}
