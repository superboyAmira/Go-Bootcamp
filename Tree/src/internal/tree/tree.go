package tree

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
	if root.isEmpty() {
		return []bool{root.HasToy}
	}
	level := 1
	var res []bool
	tmpNode := root
	for {
		res = append(res, tmpNode.HasToy)
		if level % 2 == 0 {
			tmpNode = tmpNode.Right
		} else {
			tmpNode = tmpNode.Left
		}
		level++
	}
	return res
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