func maxProduct(nums []int) int {
    if len(nums) == 0 { return -1 }
    currentMax := nums[0]
    currentMin := nums[0]
    finalMax := nums[0]
    for i := 1; i < len(nums); i++ {
        tmp := currentMax
        currentMax = max(nums[i], max(currentMax*nums[i], currentMin*nums[i]))
        currentMin = min(nums[i], min(tmp*nums[i], currentMin*nums[i]))
        if currentMax > finalMax {
            finalMax = currentMax
        }
    }
    return finalMax
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}