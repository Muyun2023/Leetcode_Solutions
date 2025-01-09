func threeSumClosest(nums []int, target int) int {
	Length := len(nums)
	sort.Ints(nums)
	var left, right, sum, diff, result int
	min := math.MaxInt
	for i := 0; i <= Length-2; i++ {
		left = i + 1
		right = Length - 1
		for left < right {
			sum = nums[left] + nums[right] + nums[i]
			diff = int(math.Abs(float64(target - sum)))
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return sum
			}
			if diff < min {
				min = diff
				result = sum
			}
		}
	}
	return result
}
