func maxSlidingWindow(nums []int, k int) []int {
	output := []int{}
	q := make([]int, 0)
	l, r := 0, 0

	for r < len(nums) {
		// pop smaller values from q
		for len(q) != 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)

		// remove left val from window
		if l > q[0] {
			q = q[1:]
		}

		if (r + 1) >= k {
			output = append(output, nums[q[0]])
			l++
		}
		r++
	}
	return output
}
