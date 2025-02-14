func combinationSum2(candidates []int, target int) [][]int {
	var backtrack func(idx int, currSum int, curr []int)
	backtrack = func(idx int, currSum int, curr []int) {
		if currSum == target {
			ans = append(ans, append([]int{}, curr...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			curr = append(curr, candidates[i])
			backtrack(i+1, currSum+candidates[i], curr)
			curr = curr[:len(curr)-1]
		}

	}
	backtrack(0, 0, curr)
	return ans
}
