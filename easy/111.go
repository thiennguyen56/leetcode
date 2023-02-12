package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type MyTreeNode struct {
	MyNode *TreeNode
	Depth  int
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	BFS := func(root *TreeNode) int {
		queue := list.New()
		queue.PushBack(MyTreeNode{
			Depth:  1,
			MyNode: root,
		})
		min := 0
		for queue.Len() > 0 {
			uAny := queue.Front()
			queue.Remove(uAny)
			uMyTreeNode := uAny.Value.(MyTreeNode)

			if uMyTreeNode.MyNode.Left != nil {
				queue.PushBack(MyTreeNode{
					MyNode: uMyTreeNode.MyNode.Left,
					Depth:  uMyTreeNode.Depth + 1,
				})
			}
			if uMyTreeNode.MyNode.Right != nil {
				queue.PushBack(MyTreeNode{
					MyNode: uMyTreeNode.MyNode.Right,
					Depth:  uMyTreeNode.Depth + 1,
				})
			}
			if uMyTreeNode.MyNode.Left == nil && uMyTreeNode.MyNode.Right == nil {
				if min == 0 || uMyTreeNode.Depth < min {
					min = uMyTreeNode.Depth
				}
			}
		}
		return min
	}
	res := BFS(root)
	return res

}
func main() {
	fmt.Println("vim-go")
}
