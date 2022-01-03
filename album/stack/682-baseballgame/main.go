package _82_baseballgame

import (
	"strconv"
)

func calPoints(ops []string) int {
	// 表达式树？
	// pervous stack
	// curSum
	stack := []int{}
	curSum := 0

	// error handling
	for _, v := range ops {
		switch v {
		case "+":
			// what if it don't have two point ?
			thisRound := stack[len(stack)-1] + stack[len(stack)-2]
			stack = append(stack, thisRound)
			curSum += thisRound
		case "D":
			p := stack[len(stack)-1]*2
			stack = append(stack, p)
			curSum += p
		case "C":
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curSum -= top
		default:
			n, _ := strconv.Atoi(v)
			curSum += n
			stack = append(stack, n)
		}
	}
	return curSum
}
