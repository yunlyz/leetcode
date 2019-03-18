package tree

import (
	"errors"
	"sync"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	nodes []*TreeNode
	sync.Mutex
}

func (s *Stack) Len() int {
	return len(s.nodes)
}

func (s *Stack) Peek() (*TreeNode, error) {
	if s.Len() < 1 {
		return nil, errors.New("out of index")
	}
	return s.nodes[s.Len()-1], nil
}

func (s *Stack) Push(val *TreeNode) bool {
	s.Lock()
	s.nodes = append(s.nodes, val)
	s.Unlock()

	return true
}

func (s *Stack) Pop() (*TreeNode, error) {
	if s.Len() < 1 {
		return nil, errors.New("out of index")
	}
	s.Lock()
	val := s.nodes[s.Len()-1]
	s.nodes = s.nodes[:s.Len()-1]
	s.Unlock()

	return val, nil
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversalByIteration(root *TreeNode) []int {
	var nodes []*TreeNode
	var order []int

	if root != nil {
		nodes = append(nodes, root)
	}

	for len(nodes) > 0 {
		// 使用堆栈来实现先入先出的特性，根据入栈的先手顺序实现前序遍历
		node := nodes[0]
		order = append(order, node.Val)
		nodes = append(nodes[:0], nodes[1:]...)
		if node.Right != nil {
			nodes = append([]*TreeNode{node.Right}, nodes...)
		}
		if node.Left != nil {
			nodes = append([]*TreeNode{node.Left}, nodes...)
		}
	}
	return order
}

func preorderTraversalByRecursive(root *TreeNode) []int {
	var result []int
	if root != nil {
		result = append(result, root.Val)
		if root.Left != nil {
			result = append(result, preorderTraversalByRecursive(root.Left)...)
		}
		if root.Right != nil {
			result = append(result, preorderTraversalByRecursive(root.Right)...)
		}
	}
	return result
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversalByIteration(root *TreeNode) []int {
	var stack Stack
	var result []int
	tree := root

	for stack.Len() > 0 || tree != nil {
		for tree != nil {
			stack.Push(tree)
			tree = tree.Left
		}
		if stack.Len() > 0 {
			t, _ := stack.Pop()
			result = append(result, t.Val)
			tree = t.Right
		}
	}
	return result
}
func inorderTraversalByRecursive(root *TreeNode) []int {
	var result []int
	if root != nil {
		if root.Left != nil {
			result = append(result, inorderTraversalByRecursive(root.Left)...)
		}
		result = append(result, root.Val)
		if root.Right != nil {
			result = append(result, inorderTraversalByRecursive(root.Right)...)
		}
	}
	return result
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversalByIterationOfAwesome(root *TreeNode) []int {
	var result []int
	var stack Stack
	var current *TreeNode

	if root != nil {
		stack.Push(root)

		for stack.Len() > 0 {
			current, _ = stack.Pop()
			if current.Left != nil {
				stack.Push(current.Left)
			}
			if current.Right != nil {
				stack.Push(current.Right)
			}
			result = append([]int{current.Val}, result...)
		}
	}
	return result
}

func postorderTraversalByRecursive(root *TreeNode) []int {
	var result []int
	if root != nil {
		if root.Left != nil {
			result = append(result, postorderTraversalByRecursive(root.Left)...)
		}
		if root.Right != nil {
			result = append(result, postorderTraversalByRecursive(root.Right)...)
		}
		result = append(result, root.Val)
	}

	return result
}
