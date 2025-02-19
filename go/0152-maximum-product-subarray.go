func maxProduct(nums []int) int {
    res, curMin, curMax := nums[0], 1, 1
    
    for i := 0; i < len(nums); i++ {
        temp := curMax * nums[i]
        curMax = max(max(nums[i] * curMax, nums[i] * curMin), nums[i])
        curMin = min(min(temp, nums[i] * curMin), nums[i])
        res = max(res, curMax)
    }
    return curMin
}

func max(a int, b int) int {
    if (a > b) {
        return a
    }
    return b
}

func min(a int, b int) int {
    if (a < b) {
        return a
    }
    return b
}

return res
