package array

import (
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	// 搜索问题
	// relative ordering
	// .
	// .
	// how to maintain relative ordering ?
	m := map[int]int{}
	for _, v := range arr2 {
		m[v] = 0
	}

	var needSort []int
	for _, v := range arr1 {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			needSort = append(needSort, v)
		}
	}
	sort.Ints(needSort)
	// fmt.Println(needSort)
	// fmt.Println(m)
	var ret []int
	for _, v := range arr2 {
		num := m[v]
		i := 1
		for i <= num {
			ret = append(ret, v)
			i++
		}
	}
	ret = append(ret, needSort...)
	return ret
}
