func maxSubArray(nums []int) int {
    result, total := nums[0], 0
    
    for i := 0; i < len(nums); i++ {
        total += nums[i]
        result = max(result, total)
        if total < 0 {
            total = 0
        }
    }
    return result
}
