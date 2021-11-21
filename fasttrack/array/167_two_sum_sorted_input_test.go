package array

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/discuss/51249/Python-different-solutions-(two-pointer-dictionary-binary-search).


// twoSum less than k
func twoSumV2(numbers []int, target int) []int {
	// sorted 二分搜索
	// two sum 的模板是？
	// two point 的模板
	// 二分查找的模板
	left:=0
	right:= len(numbers)-1

	for left < right {
		tmp:= numbers[left] + numbers[right]
		if tmp == target {
			return []int{left+1, right+1}
		}else if tmp > target {
			right --
		} else {
			left ++
		}
	}
	return []int{}
}
