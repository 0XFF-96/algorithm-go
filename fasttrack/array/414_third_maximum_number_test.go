package array

import (
	"math"
)

func thirdMax(nums []int) int {
	save := [] int {int(math.Inf(-1)),int(math.Inf(-1)),int(math.Inf(-1))}
	for _,v := range nums {
		if v != save[0] && v != save[1] && v != save[2]{
			if v > save[0]{
				save = [] int {v,save[0],save[1]}
			}else if v > save[1]{
				save = [] int {save[0],v,save[1]}
			}else if v > save[2]{
				save = [] int {save[0],save[1],v}
			}
		}
	}
	if save[2] == int(math.Inf(-1)){
		return save[0]
	}
	return save[2]
}

// follow-up question
// 如果是 find k-th 大的问题？
// 最小堆
// quick sort ！



// 不够快
// 4 ms 96.55 3.2 mb 100%
func thirdMaxV2(nums []int) int {
	save := []int{int(math.Inf(-1)),int(math.Inf(-1)),int(math.Inf(-1))}

	for _, v := range nums {
		if v != save[0] && v != save[1] && v != save[2] {
			if v > save[0] {
				save = []int{v, save[0], save[1]}
			}else if  v > save[1] {
				save = []int{save[0], v, save[1]}
			}else if v > save[2] {
				save[2] = v
			}
		}
	}
	if save[2] == int(math.Inf(-1)){
		return save[0]
	}
	return save[2]
}

// 超复杂的

func thirdMaxV3(nums []int) int {
	firstMax, secondMax, thirdMax := 0, 0, 0
	firstInit, secondInit, thirdInit := false, false, false
	for i := 0; i < len(nums); i++ {
		if (firstInit && nums[i] == firstMax) || (secondInit && nums[i] == secondMax) || (thirdInit && nums[i] == thirdMax) {
			continue
		}
		if nums[i] > firstMax || !firstInit {
			if secondInit {
				thirdMax = secondMax
				thirdInit = true
			}
			if firstInit {
				secondMax = firstMax
				secondInit = true
			}
			firstMax = nums[i]
			firstInit = true
		} else if nums[i] < firstMax && (nums[i] > secondMax || !secondInit) {
			if secondInit {
				thirdMax = secondMax
				thirdInit = true
			} else {
				secondInit = true
			}
			secondMax = nums[i]
		} else if nums[i] < secondMax && (nums[i] > thirdMax || !thirdInit) {
			thirdMax = nums[i]
			thirdInit = true
		}
	}
	if thirdInit {
		return thirdMax
	}
	return firstMax
}