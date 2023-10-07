package array

import (
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	origin := append([]int{}, nums...)
	sort.Ints(nums)

	m := make(map[int]int, len(origin))

	count := 1
	for idx,v := range nums {
		if idx != 0 && nums[idx] != nums[idx-1] {
			m[v] = count
			count++
		} else {
			m[v] = count
		}
	}

	var ret []int
	for _, v := range origin {
		ret = append(ret, m[v])
	}
	return ret
}


// 在处理情况上查一点
// dirty way to do it
//
func smallerNumbersThanCurrentV2(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	origin := append([]int{}, nums...)
	sort.Ints(nums)

	m := make(map[int]int, len(origin))

	for idx,v := range nums {
		_, ok := m[v]
		if !ok {
			m[v] = idx
		}
	}

	var ret []int
	for _, v := range origin {
		ret = append(ret, m[v])
	}
	return ret
}
