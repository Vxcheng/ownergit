package main

import (
	"fmt"

	"ownergit/learning-goSource/keyword_go/basic"
	"ownergit/learning-goSource/keyword_go/parallel" // 使用相对路径
)

func main() {
	fmt.Println("golang 25个关键字学习")
	chooseFunc("defer")
	fmt.Println("finished!!!")
}

func chooseFunc(key string) {
	switch key {
	case "select":
		parallel.Stu_select()
	case "chan":
		parallel.Stu_chan()
	case "defer":
		basic.Stu_defer()
	default:
		fmt.Println("未识别的关键字")
	}
}
