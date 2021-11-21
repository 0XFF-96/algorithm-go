package array



import "math"

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}
	// 滑动窗口的实例
	//
	//
	// two points
	// 左指针每移动一步，减去一个数
	// 右指针每移动一步，加上一个数

	maxAverage := math.MinInt64
	curSum := 0
	for i:=0;i<k; i++{
		curSum += nums[i]
	}

	//
	// leftP := 0
	for i:=4;i <len(nums); i++{

		if float64(curSum/k) > float64(maxAverage) {

		}
	}
	return 0.0
}


// 124 ms 83%, 6.8 100%
func findMaxAverageV2(nums []int, k int) float64 {
	// 滑动窗口的实例
	//
	//
	// two points
	// 左指针每移动一步，减去一个数
	// 右指针每移动一步，加上一个数
	sum := 0
	for i:=0;i<k;i++{
		sum += nums[i]
	}
	res := sum

	for i:=k;i<len(nums);i++{
		sum += nums[i] - nums[i-k]
		if sum > res {
			res = sum
		}
	}
	return float64(res)/float64(k)
}





