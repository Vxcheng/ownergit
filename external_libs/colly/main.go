package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	// 创建新的Collector
	c := colly.NewCollector()

	// 在收到响应时回调，获取整个页面HTML并保存到文件
	c.OnResponse(func(r *colly.Response) {
		fileName := "baidu.html"
		err := os.WriteFile(fileName, r.Body, 0644)
		if err != nil {
			log.Println("保存文件出错:", err)
			return
		}
		fmt.Println("页面已保存为", fileName)
	})

	// 错误处理
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("请求出错:", err)
	})

	// 访问页面
	err := c.Visit("https://www.baidu.com/")
	if err != nil {
		log.Fatal(err)
	}
}
