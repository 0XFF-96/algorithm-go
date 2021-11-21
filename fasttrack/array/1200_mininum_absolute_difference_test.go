package array

import (
	"sort"
)


// 另一种解法：https://leetcode.com/problems/minimum-absolute-difference/discuss/409575/golang-solution
// 空间利用率，还可以再提升一下。
// 现在存了很多没有必要的信息💻
func minimumAbsDifference(arr []int) [][]int {
	m := make(map[int][][]int)
	// 去掉排序
	// 在进行排序的时候，对两个元素进行比较
	// 就可以
	// O(nlogn + n)  时间复杂度
	// O(n)          空间复杂度
	sort.Ints(arr)
	minAbs := 100000000
	for i:=0;i < len(arr)-1;i++{
		tmpAbs := abs(arr[i], arr[i+1])
		if tmpAbs < minAbs {
			minAbs = tmpAbs
		}
		m[tmpAbs] = append(m[tmpAbs], []int{arr[i], arr[i+1]})
	}
	return m[minAbs]
}

//func abs(a, b int) int {
//	if a > b {
//		return a -b
//	} else {
//		return b - a
//	}
//}


