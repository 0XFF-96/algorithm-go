package array

// https://leetcode.com/problems/partition-array-into-three-parts-with-equal-sum/discuss/260895/JavaPython-3-two-O(n)-time-O(1)-space-codes-w-brief-explanation-and-analysis.
//


// 1、前缀和
// 2、后缀和
// 3、average sum / 3
// 4、


// 注意的点⚠️
// 在循环中，先加，还是先判断？取决于什么条件？
// 在结束循环的条件是什么？...
func canThreePartsEqualSum(A []int) bool {
	// 暴力解法？
	sum := 0
	for _, v := range A {
		sum += v
	}
	if sum % 3 != 0 {
		return false
	}
	cnt :=0
	average := sum / 3
	partSum := 0

	for _, v := range A {
		partSum += v
		if partSum == average {
			partSum = 0
			cnt++
		}
	}
	return cnt >= 3
}
