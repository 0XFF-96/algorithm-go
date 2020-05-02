// 能够 ac 的答案
// 但是时间成本和开销很大
// 
func canJump(nums []int) bool {
    return canJumpFromPosition(0, nums)
}

func canJumpFromPosition(position int, nums []int) bool {
    if position == len(nums) -1 {
        return true 
    }
    
    furthestJump := min(position + nums[position], len(nums)-1)
    
    for nextPos := position +1; nextPos <= furthestJump; nextPos++{
        if canJumpFromPosition(nextPos, nums) {
            return true 
        }
    }
    return false;
}

func min(i,j int) int {
    if i > j {
        return j
    } else {
        return i
    }
}
