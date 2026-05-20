package tdd

import "fmt"

/*
a(b)(

(ab()
找到剔除没成对的括号，最多一个

*/

func handleStr1(context string) (display string) {
	// 只删除一个不成对的括号
	stack := make([]int, 0, len(context))
	for i, r := range context {
		switch r {
		case '(':
			stack = append(stack, i)
		case ')':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // 匹配到一个左括号，弹出栈顶
			} else {
				return context[:i] + context[i+1:]
			}
		}
	}

	if len(stack) > 0 {
		idx := stack[len(stack)-1]
		fmt.Printf("idx: %d\n", idx)
		return context[:idx] + context[idx+1:]
	}

	return context
}

func handleStr(context string) (display string) {
	// 收集左右括号的索引
	leftSli, rightSli := make([]int, 0), make([]int, 0)
	for i, v := range context {
		if v == '(' {
			leftSli = append(leftSli, i)
		} else if v == ')' {
			rightSli = append(rightSli, i)
		}
	}

	// 左边多
	if len(leftSli) > len(rightSli) {
		return context[:leftSli[0]] + context[leftSli[0]+1:]
	} else if len(leftSli) < len(rightSli) {
		return context[:rightSli[0]] + context[rightSli[0]+1:]
	}

	return context
}
