package array
// 暴力解法
// https://leetcode.com/problems/container-with-most-water/solution/
// 两指针法
// 思路？
//


func maxArea(height []int) int {
	maxWater := 0
	for idx, v := range height {
		for i:=idx+1; i < len(height);i++{
			curWater := min(v,height[i])*(i-idx)
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


