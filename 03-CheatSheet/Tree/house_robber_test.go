package tree

import (
	"fmt"
	"testing"
)

// 这是一个用动态规划思想求解的题目
// 1、https://leetcode.com/problems/house-robber/
// 表格法
// 要求的是全局最优解法

// 极端情况 case1： 1， 2， 3， 4，1000
// case2: 1, 2, 10
// case3: 1, 2, 3, 1
// 两个比对，利用贪心法即可。
// 分解为局部性
// 1、局部有例，即全局有理


// [2, 7, 9, 3, 1]
// negetive case 1:[2, 1, 1, 2]
func TestRob(t *testing.T){

	sum := rob([]int{1,2,3})
	t.Log(sum)
}


// youtube: https://www.youtube.com/watch?v=gQzNp8gBn5k
//
func rob(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// init dp
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i:=2; i <= len(nums)-1;i++{
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
	// return max(rob[nums[:len(nums-2)]])
}


func robWrong(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	//
	if len(nums) == 1 {
		return nums[0]
	}
	sum := 0

	for i:=0; i < len(nums)-2; {
		if nums[i] > nums[i+1] {
			sum += nums[i]
		} else {
			sum += nums[i+1]
		}
		i += 2
	}
	return sum
}


// [2, 7, 9, 3, 1]
// negetive case 1:[2, 1, 1, 2]
// 证明有问题的例子2
func robWrong2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	//
	if len(nums) == 1 {
		return nums[0]
	}
	sumOdd := 0
	sumEven := 0

	// index 有问题
	for i:=1; i <= len(nums)-1; {
		sumEven += nums[i]
		i += 2
	}
	for i:=0; i <= len(nums)-1; {
		sumOdd += nums[i]
		i += 2
	}
	fmt.Println(sumOdd, sumEven)
	if sumOdd > sumEven {
		return sumOdd
	} else {
		return sumEven
	}
}