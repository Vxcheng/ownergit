package main

import (
	"fmt"
	"log"
	"time"
)

const (
	ScanInterval = 2
	layout       = "2006-01-02 15:04:05"
)

func main() {
	// covert()
	// stu_Ticker()
	// afterTimeOut()
	timeSub()

	fmt.Println("finished.")
}

func timeSub() {
	t1 := time.Now().UnixNano()
	t2 := time.Now().Unix()

	fmt.Printf("unixNano: %d, unix: %d", t1, t2)
	parseString()
	fmt.Println("finished.")
}

func parseString() {
	str := "2020-12-21T09:40:41+08:00"
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println(err)
	}

	log.Println(t.String())
}

func covert() {
	ux := time.Now().Unix()
	log.Println("ux: ", ux)
	ss := time.Unix(ux, 0).Format(layout)
	log.Println("ss: ", ss)

}

func stu_Ticker() {
	tick := time.NewTicker(time.Duration(time.Second * ScanInterval)) // 超时控制
	scanDisk(1)
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
	fmt.Printf("Unix: %d,\nUnixNano: %d,\nString: %s,\nFormat: %s\n", t.Unix(), t.UnixNano(), t.String(), t.Format(layout))
	fmt.Printf("td: %s\n", time.Date(t.Year(), t.Month(), t.Day()-1, t.Hour(), 0, 0, 0, time.UTC).String())
	fmt.Printf("===================\n")

	pastM, err := getPastTimeStamp("-1m30s")
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

func afterTimeOut() {
	log.Println("start")
	exitChan := make(chan bool)
	go func() {
		time.Sleep(time.Second * 6)
		exitChan <- true
	}()

	select {
	case s := <-exitChan:
		log.Println("signal: ", s)
	case <-time.After(time.Second * 2):
		log.Println("time out")
	}
}
