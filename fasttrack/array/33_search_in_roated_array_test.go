package array

import (
	"sort"
)

// O(N) 复杂度的算法
// https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/14425/Concise-O(log-N)-Binary-search-solution

// 相关视频：https://www.youtube.com/watch?v=7SC0hWGeyBo
//

// 另一种解法的框架：https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/501700/golang-binarysearch-solution
//

// 看看这个想法
// https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/14435/Clever-idea-making-it-simple

// 另一种条件写法：https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/544121/Golang-Binary-Search
// [4,5,6,7,0,1,2]
//0


//
// 错误典型
// 没有考虑 edge case 的情况。
// 相关题目。
// 结论1： 至少有一半部分是 sorted array, 另一半是 rated sorted array
// edge case 是什么？
func searchV1(nums []int, target int) int {
	// 二分搜索
	// 变体
	//
	l := 0
	r := len(nums) -1
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return mid
		}
		// 需要举例 几个 edge case 的情况
		// 1. 全相同
		// 2. 全递增
		// 3. 递增，中间部分相同
		// 4. 递增，中间部分相同，并从中截断
		// 判断哪一侧是 增加的
		if nums[mid] <= nums[r] {
			if target > nums[mid] {
				l = mid + 1
			} else {
				r = mid
			}
		} else {
			if target > nums[mid] {
				r = mid - 1
			} else {
				l = mid -1
			}
		}
	}

	return -1
}

// accept 的代码
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// 二分搜索
	// 变体
	//
	l := 0
	r := len(nums) -1
	for l < r {
		mid := (r + l) / 2

		if nums[mid] == target {
			return mid
		}

		// 需要举例 几个 edge case 的情况
		// 1. 全相同
		// 2. 全递增
		// 3. 递增，中间部分相同
		// 4. 递增，中间部分相同，并从中截断
		// 判断哪一侧是 增加的
		if nums[l] <= nums[mid] {
			// 这里的条件需要调换一下顺序, 为了看的人方便。
			//  target < nums[mid] && target >= nums[l]
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			}else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	// 这是为了解决哪个 edge case 而需要的代码呢？
	if nums[l] == target {
		return l
	} else {
		return -1
	}
}


// 这个好想清晰一点
func _binSearch(a []int, lo, hi, t int) int {
	r := sort.SearchInts(a[lo:hi], t)
	if r == hi-lo || a[r+lo] != t {
		return -1
	}
	return r + lo
}

func _rotatedBinSearch(a []int, lo, hi, t int) int {
	if a[lo] <= a[hi-1] {
		return _binSearch(a, lo, hi, t)
	}
	mid := lo + (hi-lo)/2
	r := _rotatedBinSearch(a, lo, mid, t)
	if r == -1 {
		r = _rotatedBinSearch(a, mid, hi, t)
	}
	return r
}

func searchV(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return _rotatedBinSearch(nums, 0, len(nums), target)
}

// https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/14435/Clever-idea-making-it-simple
//
// 比较清晰的结论
// https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/14436/Revised-Binary-Search

//public class Solution {
//public int search(int[] A, int target) {
//int lo = 0;
//int hi = A.length - 1;
//while (lo < hi) {
//int mid = (lo + hi) / 2;
//if (A[mid] == target) return mid;
//
//if (A[lo] <= A[mid]) {
//if (target >= A[lo] && target < A[mid]) {
//hi = mid - 1;
//} else {
//lo = mid + 1;
//}
//} else {
//if (target > A[mid] && target <= A[hi]) {
//lo = mid + 1;
//} else {
//hi = mid - 1;
//}
//}
//}
//return A[lo] == target ? lo : -1;
//}