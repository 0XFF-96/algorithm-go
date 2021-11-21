package array




// 这题和反转数组差不多
// 需要两次反转数组
// https://leetcode.com/articles/rotate-array/
// 一定要实现所有的算法。。。



// 这题和反转数组差不多
// 需要两次反转数组

// 边界条件在哪里？
// 极端条件
// 1个元素
//
func rotate(nums []int, k int)  {
	if len(nums) == 0 {
		return
	}

	// k = k % len(nums) -1
	//reverse(&nums, 0, k)
	//reverse(&nums, k, len(nums)-1)
	//reverse(&nums, 0, len(nums)-1)
	   k = k % (len(nums))
	   reverse(&nums, 0, len(nums)-1)
	   reverse(&nums, 0, k-1)
	   reverse(&nums, k, len(nums)-1)
	//if len(nums) == 0 || k == 0 {
	//	return
	//}
	//if len(nums) == 1 {
	//	return
	//}
	//
	//k = k % (len(nums))
	//reverse(&nums, 0, k)
	//// 这个边界条件要小心啊
	//reverse(&nums, k+1, len(nums)-1)
	//if k != len(nums)  {
	//	reverse(&nums, 0, len(nums)-1)
	//}
}


func reverse(nums *[]int, l, r int){
	if len(*nums) == 0 {
		return
	}
	numsP := *nums
	for l < r {
		numsP[l], numsP[r] = numsP[r], numsP[l]
		l++
		r--
	}
	return
}