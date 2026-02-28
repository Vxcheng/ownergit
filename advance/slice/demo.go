package slice

import "fmt"

// slice 底层数组，共享底层数组
func demo1() {
	var sl = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var parr = (*[7]int)(sl)
	var arr = *(*[7]int)(sl)
	fmt.Println(sl)  // [1 2 3 4 5 6 7]
	fmt.Println(arr) // [1 2 3 4 5 6 7]
	sl[0] = 11
	fmt.Println(sl)    // [11 2 3 4 5 6 7]
	fmt.Println(arr)   // [1 2 3 4 5 6 7]
	fmt.Println(*parr) // [11 2 3 4 5 6 7]
}

/*
s1 从 slice 索引2（闭区间）到索引5（开区间，元素真正取到索引4），长度为3，容量默认到数组结尾，为8。
s2 从 s1 的索引2（闭区间）到索引6（开区间，元素真正取到索引5），容量到索引7（开区间，真正到索引6），为5。
底层数组是可以被多个 slice 同时指向的
*/
func demo2() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}
