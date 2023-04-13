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

func levelOrderSolutionUTube(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		qLen := queue.Len()
		qRes := []int{}
		for i := 0; i < qLen; i++ {
			q := queue.Front()
			queue.Remove(q)
			qTreeNode := q.Value.(*TreeNode)
			if qTreeNode != nil {
				qRes = append(qRes, qTreeNode.Val)
				queue.PushBack(qTreeNode.Left)
				queue.PushBack(qTreeNode.Right)
			}
		}
		if len(qRes) > 0 {
			res = append(res, qRes)
		}
	}
	return res
}
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	res = append(res, []int{root.Val})
	BFS := func(root *MyTreeNode) {
		queue := list.New()
		queue.PushBack(root)
		for queue.Len() > 0 {
			uAny := queue.Front()
			queue.Remove(uAny)
			uTreeNode := uAny.Value.(*MyTreeNode)
			if uTreeNode.MyNode.Right != nil || uTreeNode.MyNode.Left != nil {
				if len(res) <= uTreeNode.Depth+1 {
					res = append(res, []int{})
				}
			}
			if uTreeNode.MyNode.Left != nil {
				res[uTreeNode.Depth+1] = append(res[uTreeNode.Depth+1], uTreeNode.MyNode.Left.Val)
				queue.PushBack(&MyTreeNode{
					MyNode: uTreeNode.MyNode.Left,
					Depth:  uTreeNode.Depth + 1,
				})
			}
			if uTreeNode.MyNode.Right != nil {
				res[uTreeNode.Depth+1] = append(res[uTreeNode.Depth+1], uTreeNode.MyNode.Right.Val)
				queue.PushBack(&MyTreeNode{
					MyNode: uTreeNode.MyNode.Right,
					Depth:  uTreeNode.Depth + 1,
				})
			}
		}
	}
	BFS(&MyTreeNode{
		MyNode: root,
		Depth:  0,
	})
	return res
}

func main() {
	fmt.Println("vim-go")
}
