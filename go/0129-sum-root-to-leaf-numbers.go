func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res []int

	dfs(root, 0, &res)
	return sum(res...)
}

func dfs(node *TreeNode, curSum int, res *[]int) {
	if node == nil {
		return
	}
	curSum = (curSum * 10) + node.Val
	if node.Left == nil && node.Right == nil {
		*res = append(*res, curSum)
	}
	dfs(node.Left, curSum, res)
	dfs(node.Right, curSum, res)
}

func sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}
