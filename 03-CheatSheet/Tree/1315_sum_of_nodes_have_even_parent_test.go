package tree

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/3sum-closest/discuss/7872/Java-solution-with-O(n2)-for-reference
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	nlen := len(nums)

	curClosest := math.MaxInt32
	for i := 0; i < nlen-2; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		start, end := i+1, nlen-1
		for start < end {
			if threeSum := nums[i] + nums[start] + nums[end]; threeSum == target {
				return threeSum
			} else if threeSum < target {
				if absV1(threeSum-target) < absV1(curClosest-target) {
					curClosest = threeSum
				}
				// 去重
				// 能够举一个具体的 case 吗？
				for start < end && nums[start] == nums[start+1] {
					start++
				}
				start++
			} else {
				if absV1(threeSum-target) < absV1(curClosest-target) {
					curClosest = threeSum
				}
				for start < end && nums[end] == nums[end-1] {
					end--
				}
				end--
			}
		}
	}
	return curClosest
}

func absV1(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}