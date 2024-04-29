package main

// 希尔排序函数
func shellSort(arr []int) {
	n := len(arr)
	gap := n / 2

	// 不断缩小步长直到为1
	for gap > 0 {
		// 对每个步长进行插入排序
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i

			// 向前比较并移动元素
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}

			// 插入元素到正确的位置
			arr[j] = temp
		}

		// 缩小步长
		gap /= 2
	}
}

// 快速排序函数
func quickSort(arr []int, low, high int) {
	if low < high {
		// 分区操作，将数组划分为两部分
		pivot := partition(arr, low, high)

		// 递归排序左半部分
		quickSort(arr, low, pivot-1)
		// 递归排序右半部分
		quickSort(arr, pivot+1, high)
	}
}

// 分区函数，选择一个基准元素，将数组分为两部分，左边部分小于基准元素，右边部分大于基准元素
func partition(arr []int, low, high int) int {
	// 选择最后一个元素作为基准元素
	pivot := arr[high]
	i := low - 1

	// 遍历数组，将小于基准元素的放到左边，大于基准元素的放到右边
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// 将基准元素放到正确的位置上
	arr[i+1], arr[high] = arr[high], arr[i+1]

	// 返回基准元素的索引
	return i + 1
}
