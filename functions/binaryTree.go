package functions

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// initializes a binary tree data structure and returns a pointer to its root node.
func InitBinaryTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	return root
}

// https://leetcode.com/problems/binary-tree-inorder-traversal/description/
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := InorderTraversal(root.Left)
	left = append(left, root.Val)
	right := InorderTraversal(root.Right)
	left = append(left, right...)
	return left
}
