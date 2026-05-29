package linked_list

import (
	"errors"
	"fmt"
)

// StackNode 链表节点
type StackNode[T any] struct {
	val  T
	next *StackNode[T]
}

// LinkedStack 链表栈
type LinkedStack[T any] struct {
	top  *StackNode[T]
	size int
}

func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{}
}

func (s *LinkedStack[T]) Push(v T) {
	node := &StackNode[T]{val: v, next: s.top}
	s.top = node
	s.size++
}

func (s *LinkedStack[T]) Pop() (T, error) {
	var zero T
	if s.top == nil {
		return zero, errors.New("stack is empty")
	}
	v := s.top.val
	s.top = s.top.next
	s.size--
	return v, nil
}

// Stack 切片栈
type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

// Push 入栈
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Pop 出栈
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.data) == 0 {
		return zero, errors.New("stack is empty")
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, nil
}

// Peek 查看栈顶
func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if len(s.data) == 0 {
		return zero, errors.New("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

// IsEmpty 是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Size 返回大小
func (s *Stack[T]) Size() int {
	return len(s.data)
}

func main_stack() {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for !stack.IsEmpty() {
		v, _ := stack.Pop()
		fmt.Print(v, " ") // 3 2 1
	}
}
