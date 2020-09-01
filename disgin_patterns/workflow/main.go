package main

import (
	"fmt"
	"log"
	pool "ownergit/disgin_patterns/workflow/pool"
	"time"
)

func main() {
	// http /router jobqueue
	var wf pool.WorkFlow
	//初始化并启动工作
	wf.StartWorkFlow(4, 20)
	for i := 0; i < 100; i++ {
		payload := pool.Payload{
			fmt.Sprintf("产品-%08d", i+1),
		}
		wJob := pool.Job{
			Payload: payload,
		}
		//添加工作
		wf.AddJob(wJob)
		//time.Sleep(time.Millisecond * 10)
	}
	wf.CloseWorkFlow()
	time.Sleep(time.Second * 4)
	log.Println("---------------------------------exit")
}
