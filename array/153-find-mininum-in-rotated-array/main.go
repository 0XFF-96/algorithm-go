package _53_find_mininum_in_rotated_array

func findMin(nums []int) int {
	// Inflection point
	lo :=0
	hi := len(nums)-1

	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < nums[hi] {
			hi = mid
		} else {
			lo = mid+1
		}
	}
	return nums[lo]
}
