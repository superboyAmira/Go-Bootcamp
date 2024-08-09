package tree

type Stack[T any] struct {
	arr []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(element T) {
	s.arr = append(s.arr, element)
}

func (s *Stack[T]) Pop() (T) {
	if len(s.arr) == 0 {
		var zero T
		return zero
	}
	topIndex := len(s.arr) - 1
	element := s.arr[topIndex]
	s.arr = s.arr[:topIndex]
	return element
}

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

	currentLevel := New[*TreeNode]()
	nextLevel := New[*TreeNode]()

	currentLevel.Push(root)

	leftStart := false

	for len(currentLevel.arr) > 0 {
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
		if len(currentLevel.arr) == 0 {
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