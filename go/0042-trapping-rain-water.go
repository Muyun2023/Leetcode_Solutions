func trap(height []int) int {
    left, right := 0, len(height) - 1
    leftMax, rightMax := height[left], height[right]
    res := 0
    
    for left < right {
        if leftMax < rightMax {
            left += 1
            leftMax = max(leftMax, height[left])
            res += leftMax - height[left]
        } else {
            right -= 1
            rightMax = max(rightMax, height[right])
            res += rightMax - height[right]
        }
    }
    return res
}

}
