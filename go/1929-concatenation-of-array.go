func getConcatenation(nums []int) []int {
  n := len(nums)
	ans := make([]int, 2*n)	
	i := 0
	for i < n {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
		i++
	}
	return ans
}
