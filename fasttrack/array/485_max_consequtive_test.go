package array


func findMaxConsecutiveOnes(nums []int) int {
	maxSeq := 0
	curSeq := 0
	for i:=0; i<len(nums); i++ {
		if nums[i] == 1 {
			curSeq++
		} else {
			curSeq = 0
		}
		if curSeq > maxSeq {
			maxSeq = curSeq
		}
	}
	return maxSeq
}


