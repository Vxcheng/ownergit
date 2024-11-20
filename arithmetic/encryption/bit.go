package encryption

import "fmt"

// pritnOddTimesNum 函数用于找出数组中出现奇数次的数字。
// 参数 arr 是一个整数数组，其中所有元素都出现偶数次，只有一个元素出现奇数次。
func printOddTimesNum(arr []int) {
	// 初始化异或结果为0，用于找出所有数字的异或结果。
	eor := 0
	// 遍历数组，通过异或操作找出出现奇数次的数字。
	for _, v := range arr {
		eor ^= v
	}

	// 通过异或结果找到最右侧的1，这一步用于将数组分为两部分，
	// 一部分包含出现奇数次的数字，另一部分不包含。
	rightOne := eor & (^eor + 1)

	// 初始化onlyOne为0，用于存储出现奇数次的数字。
	onlyOne := 0
	// 再次遍历数组，只对包含rightOne位1的数字进行异或操作，
	// 最终找出出现奇数次的数字。
	for _, v := range arr {
		if v&rightOne == 1 {
			onlyOne ^= v
		}
	}

	fmt.Println(onlyOne, eor^onlyOne)
}

func printOneOddTimesNum(arr []int) {
	// 初始化异或结果为0，用于找出所有数字的异或结果。
	eor := 0
	// 遍历数组，通过异或操作找出出现奇数次的数字。
	for _, v := range arr {
		eor ^= v
	}
	fmt.Println(eor)
}
