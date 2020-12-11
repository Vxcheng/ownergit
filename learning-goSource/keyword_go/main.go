package main

import (
	"fmt"

	"ownergit/learning-goSource/keyword_go/basic"
	"ownergit/learning-goSource/keyword_go/parallel" // 使用相对路径
)

//程序声明: package, import. (2)
//程序实体声明与定义： chan, const, func, interface, map, struct, type, var. (8)
//流程控制： break, case, continue, defer, default, else, fallthrough for, go, goto, if, rangge, return, select, switch. (15)
//内置函数：append(), cap(), close(), copy(), complex(), delete(), imag(), len() make(), new(), panic(), print(), println(), real(), recover()
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
	case "go":
		parallel.Stu_go()
	default:
		fmt.Println("未识别的关键字")
	}
}
