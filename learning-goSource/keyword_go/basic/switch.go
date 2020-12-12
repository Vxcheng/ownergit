package basic

import "log"

func Stu_switch() {
	stu1_switch()
}

func stu1_switch() {
	{
		items := []int{1, 2, 3, 4}
		for _, item := range items {
			switch item {
			case 1:
				log.Printf("value: %d\n", item)
				fallthrough
			case 5:
				log.Println("fallthrough")
			default:
				log.Println("default")
			}

		}
	}
}
