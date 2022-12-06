package slice

import "fmt"

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
