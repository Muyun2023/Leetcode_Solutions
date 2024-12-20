func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] != i+1 && (nums[i] > 0 && nums[i] <= len(nums)) && nums[i] != nums[nums[i]-1] {
			// Swap nums[i] with the number at its correct position (nums[nums[i]-1])
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	// If no missing positive number was found, return the length of the array + 1
	return len(nums) + 1
}
