package main

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	// 主函数
	arr := []int{12, 34, 54, 2, 3, 10, 1}

	fmt.Println("原始数组: ", arr)

	// 调用希尔排序函数
	shellSort(arr)

	fmt.Println("排序后的数组: ", arr)

}

func TestQuickSort(t *testing.T) {
	arr := []int{10, 7, 8, 9, 1, 5}
	n := len(arr)

	quickSort(arr, 0, n-1)
	fmt.Println("13>>1 = ", 13>>1)
	fmt.Println("排序后的数组：", arr)
}

func TestInsertSort(t *testing.T) {
	arr := []int{10, 7, 8, 9, 1, 5}
	insertSort(arr)
	fmt.Println("排序后的数组：", arr)
}
