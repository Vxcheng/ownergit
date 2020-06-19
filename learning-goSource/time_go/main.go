package main

import (
	"fmt"
	"log"
	"time"
)

const (
	ScanInterval = 5
	layout       = "2006-01-02 15:04:05"
)

func main() {
	// covert()

	printTime()
}

func covert() {
	ux := time.Now().Unix()
	log.Println("ux: ", ux)
	ss := time.Unix(ux, 0).Format(layout)
	log.Println("ss: ", ss)

}

func stu_Ticker() {
	tick := time.NewTicker(time.Duration(time.Second * ScanInterval)) // 超时控制
	i := 0
	for {
		select {
		case <-tick.C:
			i++
			scanDisk(i)
		}
	}
}
func scanDisk(i int) {
	log.Println("i=", i)
}

func printTime() {
	t := time.Now()
	fmt.Printf("Unix: %d,\nUnixNano: %d,\nString: %s,\tFormat: %s\n", t.Unix(), t.UnixNano(), t.String(), t.Format(layout))
	fmt.Printf("===================\n")

	pastM, err := getPastTimeStamp("-1h")
	if err != nil {
		log.Println("err: ", err.Error())
		return
	}
	fmt.Printf("pastM: %s\n", time.Unix(pastM, 0).Format(layout))
}

func getPastTimeStamp(str string) (int64, error) {
	d, err := time.ParseDuration(str)
	if err != nil {
		return 0, err
	}

	return time.Now().Add(d).Unix(), nil
}
