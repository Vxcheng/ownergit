package leetcode

import "container/list"

// 使用container/list实现层序遍历，计算每层的节点值之和
func levelSumWithList(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		levelSize := queue.Len()
		levelSum := 0

		for i := 0; i < levelSize; i++ {
			// 出队
			front := queue.Front()
			queue.Remove(front)
			node := front.Value.(*TreeNode)

			levelSum += node.Val

			// 子节点入队
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		result = append(result, levelSum)
	}

	return result
}
