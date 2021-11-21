package array


// 这些需要简单题，多种做法。
// 难做。
// 可以从哪几个方面优化这个函数呢？
// 以下是相关的题目。
// Longest Consecutive Sequence
//Find K-th Smallest Pair Distance
//Cells with Odd Values in a Matrix
func numPairsDivisibleBy60(time []int) int {
	// 排序。
	// pair
	// 双指针
	// 数字规律
	// 组合
	// how many pair
	// brute force two-for
	//
	var count int
	for i:=0; i < len(time); i++{
		for j:=i+1; j < len(time); j++{
			if (time[i] + time[j]) % 60 == 0 {
				count ++
			}
		}
	}
	return count
}


// https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/discuss/296138/Java-solution-from-combination-perspective-with-best-explanation
// 转变为 two-sum 的题目？
//