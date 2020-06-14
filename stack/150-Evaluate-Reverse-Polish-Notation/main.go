package _50_Evaluate_Reverse_Polish_Notation

import (
	"strconv"
)

// 后缀表达式
func evalRPN(tokens []string) int {
	var ints []int
	for _, token := range tokens{
		if token == "+"{
			n := len(ints)
			ints = append(ints[:n-2], ints[n-1]+ints[n-2])
		}else if token == "-"{
			n :=len(ints)
			ints = append(ints[:n-2], ints[n-2]-ints[n-1])
		}else if token == "/"{
			n := len(ints)
			ints = append(ints[:n-2], ints[n-2]/ints[n-1])
		}else if token == "*"{
			n := len(ints)
			ints = append(ints[:n-2], ints[n-2]*ints[n-1])
		}else{
			num, _ := strconv.Atoi(token)
			ints = append(ints, num)
		}
	}
	return ints[0]
}

