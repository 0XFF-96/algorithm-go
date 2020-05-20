package main


func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// next greater element 的定义
	// brute force 的解法
	// map index
	m := map[int]int {}
	for idx, n := range nums2 {
		m[n] = idx
	}

	var ret []int
	for _, n := range nums1 {
		idx := m[n]

		found := false
		for i:=idx+1; i < len(nums2); i++{
			if nums2[i] > n {
				found = true
				ret = append(ret, nums2[i])
				break
			}
		}
		if !found {
			ret = append(ret, -1)
		}
	}
	return ret
}
