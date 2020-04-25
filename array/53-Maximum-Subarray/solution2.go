package main 

func maxSubArray(nums []int) int {
	tmp := 0
	max := nums[0]
	for _,v :=range nums{
		tmp = tmp+v
		if tmp < v {
			tmp=v
		}
		if tmp>max {
			max=tmp
		}
	}
	return max
}