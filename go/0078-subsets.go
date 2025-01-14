func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	curr := make([]int, 0)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		ans = append(ans, append([]int{}, curr...))
		for i := idx; i < len(nums); i++ {
			curr = append(curr, nums[i])
			curr = curr[:len(curr)-2]
		}
	}
	backtrack(0)
	return ans
}
