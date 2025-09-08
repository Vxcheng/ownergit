package main

import (
	"fmt"
	"runtime"
	"time"
)

const matrixLength = 2000

func createMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
		for j := range matrix[i] {
			matrix[i][j] = i + j
		}
	}
	return matrix
}

func foo() {
	matrixA := createMatrix(matrixLength)
	matrixB := createMatrix(matrixLength)

	for i := 0; i < matrixLength; i++ {
		for j := 0; j < matrixLength; j++ {
			matrixA[i][j] = matrixA[i][j] + matrixB[i][j]
		}
	}
}

func bar() {
	matrixA := createMatrix(matrixLength)
	matrixB := createMatrix(matrixLength)

	for i := 0; i < matrixLength; i++ {
		for j := 0; j < matrixLength; j++ {
			matrixA[i][j] = matrixA[i][j] + matrixB[j][i] // 关键区别
		}
	}
}

func main() {
	// 预热
	runtime.GC()

	// 测试 foo()
	start := time.Now()
	foo()
	fooTime := time.Since(start)

	runtime.GC()

	// 测试 bar()
	start = time.Now()
	bar()
	barTime := time.Since(start)

	fmt.Printf("foo() 执行时间: %v\n", fooTime)
	fmt.Printf("bar() 执行时间: %v\n", barTime)
	fmt.Printf("bar() 比 foo() 慢: %.2f 倍\n", float64(barTime)/float64(fooTime))
}
