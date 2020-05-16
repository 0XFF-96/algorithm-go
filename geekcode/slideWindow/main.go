package slideWindow

// 滑动窗口的算法
func maxSLidingWindow(nums []int, k int) []int{
	var result []int

	if len(nums) == 0 {
		return result
	}

	if len(nums) == k {
		result = append(result, getMax(nums))
		return result
	}

	var windowSlice []int
	index := 0

	for i := k; i <= len(nums); i++ {
		windowSlice = nums[index:i]
		result = append(result, getMax(windowSlice))
		index++
	}
	return result
}

func getMax(windowSlice []int) int {
	max := windowSlice[0]
	for i := 0; i < len(windowSlice); i++{
		if windowSlice[i] >= max{
			max = windowSlice[i]
		}
	}
	return max
}

func maxSlidingWindow(nums []int, k int) []int{
	result := []int{}

	if len(nums) == 0 {
		return result
	}

	// 窗口如何移动
	// 最大值如何被获取
	window := []int{}
	for i, x := range nums {
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) != 0 && nums[window[len(window)-1]] <= x {
			window = window[1:]
		}
		window = append(window, i)
		if i >= k -1 {
			result = append(result, nums[window[0]])
		}
	}
	return result
}
