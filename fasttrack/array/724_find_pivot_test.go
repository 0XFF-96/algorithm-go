package array

import (
	"testing"
)

func TestFindPivot(t *testing.T){
	sum := pivotIndex([]int{1, 7, 3, 6, 5, 6})
	t.Log(sum)
}


// 没有 AC 的答案
// case : [-1,-1,-1,-1,-1,0]
// 查找的方法.
func pivotIndex(nums []int) int {
	if len(nums) == 0  {
		return -1
	}
	left := 0
	right := len(nums)-1
	leftSum := 0
	rightSum := 0
	for left < right {
		if leftSum <= rightSum {
			leftSum += nums[left]
			left ++
		} else {
			rightSum += nums[right]
			right--
		}
	}

	if leftSum == rightSum {
		return left
	} else {
		return -1
	}
	// fmt.Println(rightSum, leftSum, right, left)
	//if leftSum - nums[left] == rightSum {
	//	return left
	//} else {
	//	return -1
	//}
}


// good students.
// If pivot index is 0 or nums.length - 1,why its leftsum or
// rightsum is 0,there are no numbers here.
// I mean the sum beyond the boundary is 0?I am a newbie,
// if you explain it to me, thank you.


// 20 ms 6 MB   80% faster 100% saved memory
func pivotIndexV2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	totalSum := 0
	leftSum := 0
	for _, n := range nums {
		totalSum += n
	}

	for i:= 0;i < len(nums); i++{
		if (leftSum == totalSum - leftSum - nums[i]) {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}