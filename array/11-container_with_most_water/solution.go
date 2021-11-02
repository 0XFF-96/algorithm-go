package main 

// 非常暴力的解法。
func maxArea(height []int) int {
	maxWater := 0
	for idx, v := range height {
		for i:=idx+1; i < len(height);i++{
			curWater := min(v, height[i]) *   (i-idx)
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
