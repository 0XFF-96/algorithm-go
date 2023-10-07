package array



// 思路
// 1、不变性
// 2、等差数列的性质 = 和有一个固定的公式...
// 3、
// https://leetcode.com/problems/missing-number/solution/
// 里面有几种不同的实现方法，都值得好好看一下...
// missing number:https://www.youtube.com/watch?v=z0p_FBatGWM
// 有什么故事？
func missingNumber(nums []int) int {

	// 哨兵节点..
	// 假设就是最后一个缺失了..
	miss := len(nums)
	for i:=0;i<len(nums);i++{
		miss = miss ^ i ^ nums[i]
	}
	return miss
}



