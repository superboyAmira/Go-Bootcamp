package tree

import (
	"log"
	"testing"
)

func TestNewTree(t *testing.T) {
	root := NewTree(true)
	if root == nil {
		t.Error("NewTree() returned nil, expected a new TreeNode")
	}
	if !root.HasToy {
		t.Error("NewTree(true) should have HasToy set to true")
	}
}

func TestIsEmpty(t *testing.T) {
	root := NewTree(false)
	if !root.isEmpty() {
		t.Error("isEmpty() should return true for a tree with no children")
	}

	root.Left = NewTree(false)
	if root.isEmpty() {
		t.Error("isEmpty() should return false for a tree with children")
	}
}

func TestGetValueSubTree(t *testing.T) {
	root := NewTree(true)
	root.Left = NewTree(true)
	root.Right = NewTree(false)
	root.Left.Left = NewTree(true)

	var value int64
	getValueSubTree(root, &value)

	if value != 3 {
		t.Errorf("getValueSubTree() returned %d, expected 3", value)
	}
}

func TestAreToysBalanced(t *testing.T) {
	root := NewTree(true)
	if !root.areToysBalanced() {
		t.Error("areToysBalanced() should return true for a single node tree")
	}

	root.Left = NewTree(true)
	root.Right = NewTree(false)

	if root.areToysBalanced() {
		t.Error("areToysBalanced() should return false when left and right subtrees are not balanced")
	}

	root.Right.HasToy = true

	if !root.areToysBalanced() {
		t.Error("areToysBalanced() should return true when left and right subtrees are balanced")
	}

	root.Left.Left = NewTree(false)
	root.Left.Right = NewTree(false)
	root.Right.Left = NewTree(false)
	root.Right.Right = NewTree(false)

	if !root.areToysBalanced() {
		t.Error("areToysBalanced() should return true when left and right subtrees are balanced")
	}

	root.Left.Left.HasToy = true

	if root.areToysBalanced() {
		t.Error("areToysBalanced() should return false when left and right subtrees are not balanced")
	}
}

func TestUnrollGarland(t *testing.T) {
	root := &TreeNode{
		HasToy: true,
		Left: &TreeNode{
			HasToy: true,
			Left: &TreeNode{
				HasToy: true,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				HasToy: false,
				Left: nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			HasToy: false,
			Left: &TreeNode{
				HasToy: true,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				HasToy: true,
				Left: nil,
				Right: nil,
			},
		},
	}

	res := unrollGarland(root)
	log.Print(res)
	origin := []bool{ true, true, false, true, true, false, true }
	if len(res) != len(origin) {
		t.Errorf("Failed test: result = %v, origin = %v", res, origin)
	}
	for i, elem := range res {
		if elem != origin[i] {
			t.Errorf("Failed test: result = %v, origin = %v", res, origin)
			return
		}
	}


	root = &TreeNode{
		HasToy: true,
		Left: &TreeNode{
			HasToy: false,
			Left: &TreeNode{
				HasToy: true,
				Left: &TreeNode{
					HasToy: true,
					Left: &TreeNode{
						HasToy: true,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						HasToy: false,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					HasToy: true,
					Left: &TreeNode{
						HasToy: true,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						HasToy: false,
						Left:  nil,
						Right: nil,
					},
				},
			},
			Right: &TreeNode{
				HasToy: false,
				Left: &TreeNode{
					HasToy: true,
					Left: &TreeNode{
						HasToy: true,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						HasToy: false,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					HasToy: true,
					Left: &TreeNode{
						HasToy: true,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						HasToy: true,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		Right: &TreeNode{
			HasToy: true,
			Left: &TreeNode{
				HasToy: true,
				Left: &TreeNode{
					HasToy: false,
					Left: &TreeNode{
						HasToy: false,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						HasToy: true,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					HasToy: true,
					Left: &TreeNode{
						HasToy: false,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						HasToy: true,
						Left:  nil,
						Right: nil,
					},
				},
			},
			Right: &TreeNode{
				HasToy: false,
				Left: &TreeNode{
					HasToy: false,
					Left: &TreeNode{
						HasToy: true,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						HasToy: true,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					HasToy: true,
					Left: &TreeNode{
						HasToy: false,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						HasToy: true,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}
	res = unrollGarland(root)
	log.Println(res)
	origin = []bool{true, false, true, false, true, false, true, true, true, true, true, false, true, false, true, true, false, true, true, true, false, true, false, true, true, false, true, false, true, false, true}
	if len(res) != len(origin) {
		t.Errorf("Failed test: result = %v, origin = %v", res, origin)
	}
	for i, elem := range res {
		if elem != origin[i] {
			t.Errorf("Failed test: result = %v, origin = %v", res, origin)
		}
	}
}