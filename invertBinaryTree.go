/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	var left *TreeNode = invertTree(root.Right)
	var right *TreeNode = invertTree(root.Left)

	root.Left = left
	root.Right = right

	return root
}