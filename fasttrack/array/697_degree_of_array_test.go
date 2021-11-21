package array

import (
	"math"
)

// https://leetcode.com/problems/degree-of-an-array/solution/
// 这题不懂，可以继续看一下相关的东西。
//

// 一个开始我是不知道，题目要干什么的
// https://www.youtube.com/watch?v=BY6rz-DIV28
// 懂了题目在干什么
//

// 有几种不同的 case
// 需要考虑到的情况是
// https://leetcode.com/problems/degree-of-an-array/solution/
//

// 看看题目如何实现？
//

// 构造几个测试用例看看。
// 能否通过。
// 如何构造测试用例。
func findShortestSubArray(nums []int) int {
	tmp := make(map[int][]int)
	degree := 0
	subArr := 50000

	for idx, n := range nums {
		// 这里为什么不需要初始化？
		tmp[n] = append(tmp[n], idx)
		if len(tmp[n]) > degree {
			degree = len(tmp[n])
		}
	}

	for k, v := range tmp {
		if len(v) == degree {

			// 这里为什么需要加 1 ？
			l := v[len(tmp[k])-1] - v[0] + 1
			if l < subArr {
				subArr = l
			}
		}
	}
	return subArr
}


func findShortestSubArrayV2(nums []int) int {
	left := make(map[int]int)
	right := make(map[int]int)
	count := make(map[int]int)
	max := 0
	for i:=0 ; i < len(nums); i++ {
		x := nums[i]
		if _, ok := left[x]; !ok {
			left[x] = i
		}
		right[x] = i
		val, ok := count[x];
		if ok {
			count[x] = 0
		}
		count[x] = val + 1
		if max < count[x] {
			max = count[x]
		}
	}

	ans := len(nums)
	if ans == 1 {
		return 1
	}
	for k, v := range count {
		if v == max {
			newans := (right[k] - left[k]) + 1
			if newans < ans {
				ans = newans
			}
		}
	}
	return ans
}


//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}


func findShortestSubArrayV3(nums []int) int {
	hashMapStart := make(map[int]int, len(nums))
	hashMapEnd := make(map[int]int, len(nums))
	hashMapCnts := make(map[int]int, len(nums))
	for i:=0; i<len(nums); i++ {
		if _, ok := hashMapStart[nums[i]]; !ok {
			hashMapStart[nums[i]] = i
		}
		hashMapCnts[nums[i]]++
		hashMapEnd[nums[i]] = i
	}

	maxDegree := 0
	for _, v := range hashMapCnts {
		maxDegree = max(maxDegree, v)
	}

	cnts := math.MaxInt32
	for i:=0; i<len(nums); i++ {
		if maxDegree == hashMapCnts[nums[i]] {
			cnts = min(cnts, hashMapEnd[nums[i]]-hashMapStart[nums[i]]+1)
		}
	}
	return cnts
}