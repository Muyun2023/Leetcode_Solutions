func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] != i+1 && (nums[i] > 0 && nums[i] <= len(nums)) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
、	return len(nums) + 1
}
