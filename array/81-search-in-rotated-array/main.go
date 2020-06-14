package _1_search_in_rotated_array



// This is a follow up problem to Search in Rotated Sorted Array, where nums may contain duplicates.
// Would this affect the run-time complexity? How and why?


// 复用之前的代码， 有两个 test case 不通过。
// [1, 3, 1, 1, 1, 1]
// [1, 3]







// 其他情况。
//
// AC 的情况

func search(nums []int, target int) bool {
	pivot := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			pivot = i
			break
		}
	}
	if lp := binarySearch(nums[0:pivot], target); lp != -1 {
		return true
	}
	if rp := binarySearch(nums[pivot:], target); rp != -1 {
		return true
	}
	return false
}

func binarySearch(n []int, target int) int {
	l := 0
	r := len(n)-1
	for ; l <= r; {
		mid := (l+r)/2
		switch {
		case n[mid] == target:
			return mid
		case n[mid] < target:
			l = mid+1
		case n[mid] > target:
			r = mid-1
		}
	}
	return -1
}