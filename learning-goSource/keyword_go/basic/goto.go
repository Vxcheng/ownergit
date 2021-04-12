package basic

import (
	"log"
	"time"
)

func Stu_Goto() {
	log.Println("learn goto")
	stu1_goto()
	log.Println("learn goto finished")
}

func stu1_goto() {
tag:
	for i, ele := range []int{1, 2, 3} {
		log.Printf("%d\n", ele)
		if i == 1 {
			break tag
		}
	}
}

func stu2_goto() {
tag:
	for i, ele := range []int{1, 2, 3} {
		log.Printf("%d\n", ele)
		if i == 1 {
			continue tag
		}
	}
}

func stu3_goto() {

	for i, ele := range []int{1, 2, 3} {
		log.Printf("%d\n", ele)
		if i == 1 {
			goto tag
		}
	}
tag:
}

func stu4_goto() {
tag:
	for i, ele := range []int{1, 2, 3} {
		log.Printf("%d\n", ele)
		if i == 1 {
			time.Sleep(time.Second)
			goto tag
		}
	}

}
