package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

// 代表一个任务
type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time // expr.Next(now)
}

func main() {
	// 定时任务字典， key: 任务的名字, value 任务对象
	scheduleTable := make(map[string]*CronJob)
	now := time.Now()

	// 定义定时任务以每 5s 执行一次
	// MustParse 如果遇到解析 contab 错误时会直接抛出 panic ，不会像 Parse 一样返回一个错误
	expr := cronexpr.MustParse("*/5 * * * * * *")
	cronJob := &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	// 任务注册到调度表
	scheduleTable["job1"] = cronJob

	// 定义定时任务以每 3s 执行一次
	expr = cronexpr.MustParse("*/3 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	// 任务注册到调度表
	scheduleTable["job2"] = cronJob

	// 定时检查一下任务调度表
	for {
		now := time.Now()

		for jobName, cronJob := range scheduleTable {
			// 判断是否到期，当前时间等于定时任务的下次执行时间，或者当前时间大于任务的定时时间
			if now.Equal(cronJob.nextTime) || now.After(cronJob.nextTime) {
				// 启动一个协程, 执行这个任务
				go func(jobName string) {
					fmt.Println("执行:", jobName)
				}(jobName)

				// 计算下一次调度时间
				cronJob.nextTime = cronJob.expr.Next(now)
				fmt.Println(jobName, "下次执行时间:", cronJob.nextTime)
			}
		}

		// 等待 1s，减少 CPU 消耗
		t := <-time.NewTimer(1 * time.Second).C
		fmt.Println(t)
	}

}
