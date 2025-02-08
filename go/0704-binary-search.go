func search(nums []int, target int) int {
	left := 0
	right := len(nums) 
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle
		} else {
			right = middle - 1
		}
	}
	return -1

}
