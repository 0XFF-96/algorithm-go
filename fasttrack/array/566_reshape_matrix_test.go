package array


// 不知道怎么做
// 看看这个老哥的出事话方法
// https://leetcode.com/problems/reshape-the-matrix/discuss/292585/Golang-easy-understand
//

// mapping original solution
// https://leetcode.com/problems/reshape-the-matrix/discuss/292585/Golang-easy-understand
//


// AC 的答案
// 有几种不同的方法做...
//
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || r*c != len(nums) * len(nums[0]) {
		return nums
	}
	ret := make([][]int, r)
	for i:=0;i<r;i++{
		ret[i] = make([]int, c)
	}

	count := 0
	for i:=0 ; i <len(nums); i++{
		for j := 0;j < len(nums[0]);j++{
			ret[count/c][count % c] = nums[i][j]
			count++
		}
	}
	return ret
}