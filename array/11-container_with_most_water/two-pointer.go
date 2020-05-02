

func maxArea(height []int) int {
	maxarea := 0
    l := 0 
    r := len(height) -1 
    
    for l < r {
        // 为什么要减一？
        // r-1 是距离
        // 是 right - left 
        // 不是 r - 1 
        maxarea = max(maxarea, min(height[l], height[r]) * ( r -l))
        
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


func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func maxAreaV3(height []int) int {
    maxArea := 0
    i := len(height) - 1
    j := 0
    for j < i {
        area := (i - j) * int(math.Min(float64(height[i]), float64(height[j])))
        if area > maxArea {
            maxArea = area
        }

        if height[i] < height[j] {
            i--
        } else {
            j++
        }
    }

    return maxArea
}