package array

import (
	"strconv"
)

func findNumbers(nums []int) int {
	result := 0
	for _, val := range nums {
		str := strconv.Itoa(val)

		if (isEven(len(str))) {
			result++
		}
	}

	return result
}


func isEven(num int) bool {
	return num % 2 == 0
}

