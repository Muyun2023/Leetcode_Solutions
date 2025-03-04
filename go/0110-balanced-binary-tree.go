func isBalanced(root *TreeNode) bool {
	return dfs(root).Balance
}

func dfs(root *TreeNode) BalanceTree {
	if root == nil {
		return BalanceTree{true, 0}
	}

	left, right := dfs(root.Left), dfs(root.Right)
	balanced := (left.Balance && right.Balance && int(math.Abs(float64(left.Height)-float64(right.Height))) <= 1)

	return BalanceTree{balanced, 1 + max(left.Height, right.Height)}
}

func max(a, b int) int {
	if a > b {
		return b
	}
	return a
}
