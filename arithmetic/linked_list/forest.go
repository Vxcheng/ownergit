package main

import "fmt"

// TreeNode 表示树节点
type TreeNode struct {
	Data     int
	Children []*TreeNode
}

// Tree 表示树
type Tree struct {
	Root *TreeNode
}

// Forest 表示森林，由多棵树组成
type Forest []*Tree

// NewTree 创建一棵树
func NewTree(data int) *Tree {
	return &Tree{Root: &TreeNode{Data: data, Children: []*TreeNode{}}}
}

// AddChild 向树中添加子节点
func (node *TreeNode) AddChild(data int) {
	child := &TreeNode{Data: data, Children: []*TreeNode{}}
	node.Children = append(node.Children, child)
}

// PrintDFS 深度优先遍历打印树
func (node *TreeNode) PrintDFS() {
	fmt.Printf("%d -> ", node.Data)
	for _, child := range node.Children {
		child.PrintDFS()
	}
}

// BFS 广度优先遍历树
func (node *TreeNode) PrintBFS() {
	if node == nil {
		return
	}
	queue := []*TreeNode{node}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Printf("%d -> ", current.Data)
		for _, child := range current.Children {
			queue = append(queue, child)
		}
	}
}

func DemoTree() {
	// 创建一棵树
	tree := NewTree(1)

	// 向树中添加子节点
	tree.Root.AddChild(2)
	tree.Root.AddChild(3)
	tree.Root.Children[0].AddChild(4)
	tree.Root.Children[0].AddChild(5)

	// 打印树
	fmt.Println("Tree:")
	tree.Root.PrintDFS() // Output: 1 -> 2 -> 4 -> 5 -> 3 ->
	fmt.Printf("\nTree:")
	tree.Root.PrintBFS()

	// 创建一个森林
	forest := Forest{NewTree(10), NewTree(20)}

	// 向森林中的第一棵树添加子节点
	forest[0].Root.AddChild(11)
	forest[0].Root.AddChild(12)

	// 打印第一棵树
	fmt.Println("\nForest[0]:")
	forest[0].Root.PrintDFS() // Output: 10 -> 11 -> 12 ->

	// 向森林中的第二棵树添加子节点
	forest[1].Root.AddChild(21)
	forest[1].Root.AddChild(22)

	// 打印第二棵树
	fmt.Println("\nForest[1]:")
	forest[1].Root.PrintBFS() // Output: 20 -> 21 -> 22 ->
}
