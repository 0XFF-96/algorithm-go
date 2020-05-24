package _5_sort_colors

func sortColors(nums []int) {
	counters := make([]int, 3)
	for _, val := range nums {
		counters[val]++
	}

	index := 0
	for val, cnt := range counters {
		for i := 0; i < cnt; i++ {
			nums[index] = val
			index++
		}
	}
}

