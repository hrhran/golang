package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	var h func(*TreeNode)
	h = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		if cur.Left != nil && cur.Left.Left == nil && cur.Left.Right == nil {
			sum += cur.Left.Val
		}
		h(cur.Left)
		h(cur.Right)
	}
	h(root)
	return sum
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	result := sumOfLeftLeaves(root)
	fmt.Println("Sum of left leaves:", result)
}
