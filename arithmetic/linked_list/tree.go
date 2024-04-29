package main

import "fmt"

type BiTree interface {
	Inorder(tree *Node) /*中序遍历*/

	Preorder(tree *Node) /*先序遍历*/

	Postorder(tree *Node) /*后续遍历*/

	Bforder(tree *Node) /*广度优先遍历*/

	Dforder(tree *Node) /*深度优先遍历*/
}

/*二叉树树的每一个节点数据结构*/
type Node struct {
	Data interface{} /*树的数据*/

	Left *Node /*树的左节点*/

	Right *Node /*树的右节点*/
}

func (n Node) Inorder(tree *Node) {
	if tree == nil {
		return
	}
	n.Inorder(tree.Left)
	fmt.Print(tree.Data, " ")
	n.Inorder(tree.Right)
}

func (n Node) Preorder(tree *Node) {
	if tree == nil {
		return
	}
	fmt.Print(tree.Data, " ")
	n.Preorder(tree.Left)
	n.Preorder(tree.Right)
}

func (n Node) Postorder(tree *Node) {
	if tree == nil {
		return
	}
	n.Postorder(tree.Left)
	n.Postorder(tree.Right)
	fmt.Print(tree.Data, " ")
}
