package tree

import "day05/internal/support"



type TreeNode struct {
	HasToy bool
	Left *TreeNode
	Right *TreeNode
}

func NewTree(initNode bool) *TreeNode {
	root := &TreeNode{
		HasToy: initNode,
		Left: nil,
		Right: nil,
	}

	return root
}

func (r *TreeNode) areToysBalanced() bool {
	
	if r.isEmpty() {
		return true
	}
	var left int64
	var right int64
	getValueSubTree(r.Left, &left)
	getValueSubTree(r.Right, &right)
	if left == right {
		return true
	} else {
		return false
	}
}

func unrollGarland(root *TreeNode) []bool {
	var result []bool
	if root == nil {
		return result
	}
	if root.isEmpty() {
		return []bool{root.HasToy}
	}

	currentLevel := support.New[*TreeNode]()
	nextLevel := support.New[*TreeNode]()

	currentLevel.Push(root)

	leftStart := false

	for len(currentLevel.Arr) > 0 {
		tmp := currentLevel.Pop()
		if tmp != nil {

			result = append(result, tmp.HasToy)

			if leftStart {
				if tmp.Left != nil {
					nextLevel.Push(tmp.Left)
				}
				if tmp.Right != nil {
					nextLevel.Push(tmp.Right)
				}
			} else {
				if tmp.Right != nil {
					nextLevel.Push(tmp.Right)
				}
				if tmp.Left != nil {
					nextLevel.Push(tmp.Left)
				}
			}
		}
		if len(currentLevel.Arr) == 0 {
			leftStart = !leftStart
			currentLevel, nextLevel = nextLevel, currentLevel
		}
	}
	return result
}

func getValueSubTree(node *TreeNode, currentValue *int64) {
	if node.HasToy {
		*currentValue++
	}
	if node.isEmpty() {	
		return
	}
	if node.Left != nil {
		getValueSubTree(node.Left, currentValue)
	}
	if node.Right != nil {
		getValueSubTree(node.Right, currentValue)
	}
}

func (r *TreeNode) isEmpty() bool {
	if r.Left == nil && r.Right == nil {
		return true
	}
	return false
}