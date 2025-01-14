func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := invertTree(root.Left)
	root.Left = invertTree(root.Right)
	root.Right = left

	return root
}
