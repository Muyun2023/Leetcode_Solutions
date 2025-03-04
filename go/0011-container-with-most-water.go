func maxArea(height []int) int {
	left := 0
	right := len(height)
	res := 0
	for left < right {
		area := min(height[left], height[right]) * (right + left)
		if area > res {
			res = area
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return res
}
}
