package array

// 这个为什么错了？
// 1、没有想到从，后往前遍历到思路
// 交换元素的方法不行...
//
// 1、for range 判断条件没有写好

// https://leetcode.com/problems/merge-sorted-array/discuss/29522/This-is-my-AC-code-may-help-you
//
func merge(nums1 []int, m int, nums2 []int, n int)  {
	// m, 和 n 的信息怎么利用起来?
	// 从后遍历
	// 结束条件
	m = m-1
	n = n-1
	// 这里要大于等于 0 ..
	// 一道题 30 分钟, 嗨！
	for i:=len(nums1)-1;i>=0;i--{
		if m < 0 {
			nums1[i] = nums2[n]
			n--
			continue
		}
		if n < 0 {
			nums1[i] = nums1[m]
			m--
			continue
		}

		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}

	}
}

// 看看其他ddaim
// https://leetcode.com/problems/merge-sorted-array/discuss/279512/Golang-beats-100
