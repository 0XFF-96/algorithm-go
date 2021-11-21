package array

import (
	"sort"
)

// 没有能够 AC
// Shame !
// Shame !
// https://leetcode.com/problems/3sum/
func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	var ret [][]int
	for idx, v := range nums{
		// fix number v
		vmap := map[int]struct{}{}
		for i:= idx+1;i <len(nums);i++{
			tmp := -v - nums[i]
			if _,ok:= vmap[tmp];ok {
				ret = append(ret, []int{tmp, v, nums[i]})
			} else {
				vmap[tmp] = struct{}{}
			}
		}
	}
	return ret
}


// 基本正确，但是没有AC
// 因为没有去重
// 记录哪些值很重要!
func threeSumV2(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	// 没有去重
	// 需要去重的方法
	var ret [][]int
	for idx, v := range nums{
		// fix number v
		vmap := map[int]int{}
		for i:= idx+1;i <len(nums);i++{
			tmp := -v - nums[i]
			if _,ok:= vmap[tmp];ok {
				ret = append(ret, []int{tmp, v, nums[i]})
			} else {
				vmap[nums[i]] = i
			}
		}
	}
	return ret
}


// 排序去重..
// https://leetcode.com/problems/3sum/discuss/7513/golang-29ms-a-little-optimize
// https://leetcode.com/problems/3sum/discuss/7402/Share-my-AC-C%2B%2B-solution-around-50ms-O(N*N)-with-explanation-and-comments
// 与 two sum 的不同之处在于 XXX
// 相关知识...
// 看看这相关答案
// 利用 sort 去重
// 利用 map 去重
// 能否举一个具体的例子看看？
// 真正的一个例子
// 最后实现了
func threeSumV3(nums []int) [][]int {
	sort.Ints(nums)

	var ret [][]int
	for i:=0;i < len(nums);i++{
		if i>0 && nums[i] == nums[i-1] {
			continue
		}
		left := i+1
		right := len(nums) -1
		for left < right {
			val := nums[i] + nums[left] + nums[right]
			switch  {
			case val == 0:
				// 为什么这里的 left + 1 不会超？
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				left, right = left+1, right-1

			case val > 0:
				right--
			case val < 0:
				left++
			}
		}
	}

	return ret
}