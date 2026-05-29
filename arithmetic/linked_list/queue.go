package linked_list

import "errors"

// Queue 切片队列
type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

// Enqueue 入队
func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

// Dequeue 出队
func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if len(q.data) == 0 {
		return zero, errors.New("queue is empty")
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v, nil
}

// Front 查看队首
func (q *Queue[T]) Front() (T, error) {
	var zero T
	if len(q.data) == 0 {
		return zero, errors.New("queue is empty")
	}
	return q.data[0], nil
}

// IsEmpty 是否为空
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Size 返回大小
func (q *Queue[T]) Size() int {
	return len(q.data)
}

// CircularQueue 循环队列
type CircularQueue[T any] struct {
	data []T
	head int
	tail int
	size int
	cap  int
}

func NewCircularQueue[T any](capacity int) *CircularQueue[T] {
	return &CircularQueue[T]{
		data: make([]T, capacity),
		head: 0,
		tail: 0,
		size: 0,
		cap:  capacity,
	}
}

func (q *CircularQueue[T]) Enqueue(v T) error {
	if q.size == q.cap {
		return errors.New("queue is full")
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.cap
	q.size++
	return nil
}

func (q *CircularQueue[T]) Dequeue() (T, error) {
	var zero T
	if q.size == 0 {
		return zero, errors.New("queue is empty")
	}
	v := q.data[q.head]
	q.head = (q.head + 1) % q.cap
	q.size--
	return v, nil
}

// DequeNode 双向链表节点
type DequeNode[T any] struct {
	val  T
	prev *DequeNode[T]
	next *DequeNode[T]
}

// LinkedDeque 链表双向队列
type LinkedDeque[T any] struct {
	head *DequeNode[T]
	tail *DequeNode[T]
	size int
}

func NewLinkedDeque[T any]() *LinkedDeque[T] {
	return &LinkedDeque[T]{}
}

func (d *LinkedDeque[T]) PushFront(v T) {
	node := &DequeNode[T]{val: v}
	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		node.next = d.head
		d.head.prev = node
		d.head = node
	}
	d.size++
}

func (d *LinkedDeque[T]) PushBack(v T) {
	node := &DequeNode[T]{val: v}
	if d.tail == nil {
		d.head = node
		d.tail = node
	} else {
		node.prev = d.tail
		d.tail.next = node
		d.tail = node
	}
	d.size++
}

func (d *LinkedDeque[T]) PopFront() (T, error) {
	var zero T
	if d.head == nil {
		return zero, errors.New("deque is empty")
	}
	v := d.head.val
	d.head = d.head.next
	if d.head != nil {
		d.head.prev = nil
	} else {
		d.tail = nil
	}
	d.size--
	return v, nil
}

func (d *LinkedDeque[T]) PopBack() (T, error) {
	var zero T
	if d.tail == nil {
		return zero, errors.New("deque is empty")
	}
	v := d.tail.val
	d.tail = d.tail.prev
	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.head = nil
	}
	d.size--
	return v, nil
}
