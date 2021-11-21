package array

import "sort"

// 64ms, 6.5mb 73% 100 %
// 这里刚开始有点看不懂题目点意思


// 还有没其他方式，
// 可以优化？
// sort 的时候，进行算数。节省一个 loop
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i:=0; i < len(nums);{
		sum += nums[i]
		i += 2
	}
	return sum
}

// 看思考🤔
// 如何进行证明的
// https://leetcode.com/problems/array-partition-i/discuss/102170/Java-Solution-Sorting.-And-rough-proof-of-algorithm.
//

// https://leetcode.com/problems/array-partition-i/discuss/102201/Python-solution-with-detailed-explanation