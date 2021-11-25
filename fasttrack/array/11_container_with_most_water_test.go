package array
// 暴力解法
// https://leetcode.com/problems/container-with-most-water/solution/
// 两指针法, 能够使用图表说明代码。
// 思路？
// 
func maxArea(height []int) int {
	maxWater := 0

	// 两个 for 循环
	// 能否 early-break , 提前返回
	for idx, v := range height {
		for i:=idx+1; i < len(height);i++{
			curWater := min(v, height[i])*(i-idx)
			if curWater > maxWater {
				maxWater = curWater
			}
		}
	}
	return maxWater
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// brute force 
// how to prove it is true !
func maxAreaV2(height []int) int {
	maxarea := 0
    l := 0 
    r := len(height) -1 
    
    for l < r {
        // 为什么要减一？
        // r-1 是距离
        // 是 right - left 
        // 不是 r - 1 
        maxarea = max(maxarea, min(height[l], height[r]) * ( r -l))
        
		// how to move the pointer 
        if height[l] < height[r] {
            l++
        } else {
            r --
        }
        
    }
	return maxarea
}

func max(i, j int) int {
    if i > j {
        return i
    } else {
        return j
    }
} 