package array

// find ----> search 搜索的题目
// 思路：二分查找 + hashtable 的相关优化

// 不会做
// 看了相关答案，
// https://leetcode.com/problems/k-diff-pairs-in-an-array/discuss/100098/Java-O(n)-solution-one-Hashmap-easy-to-understand
// 1、
// 2、
// 3、

// 这里面，需要去重之后，就可以了..
func findPairs(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	cur := 0
	for idx, v := range nums {
		for i:=idx+1; i < len(nums); i++{
			if abs(v, nums[i]) == k {
				cur++
			}
		}
	}
	return cur
}



func findPairsV2(nums []int, K int) int {
	if K < 0 {
		return 0
	}
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	var out int
	for k, _ := range m {
		v2 := m[k+K]
		// 这个条件是什么意思？
		if (K == 0 && v2 == 1) || v2 == 0 {
			continue
		}
		out++
	}
	return out
}



func abs(a, b int) int {
	if a > b {
		return a -b
	} else {
		return b -a
	}
}
