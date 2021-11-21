package array



// AC: 一次
// 看看参考答案如何写的.
func searchRange(nums []int, target int) []int {
	if len(nums) ==0 {
		return []int{-1, -1}
	}
	// 二分搜索
	// 如果有相同元素，向左边移动。
	l := 0
	r := len(nums) -1
	//  var ret []int
	for l < r {
		mid := (l +r)/2
		// == 的情况怎么办？
		// 有重复元素怎么办？
		if nums[mid] < target {
			l = mid +1
		} else {
			r = mid
		}
	}
	// 5， 7， 7， 8， 8， 10 tareget = 8 or target = 7 ？
	// 5, 7, 7, 7, 8, 10 target = 7 的时候，怎么办？
	if nums[l] != target {
		return []int{-1, -1}
	}

	// 2, 2
	rightMost := r
	for i:=r+1; i < len(nums);i++{
		if nums[i] == target {
			rightMost++
		}
	}
	// 线性查找不好
	return []int{r, rightMost}
}
