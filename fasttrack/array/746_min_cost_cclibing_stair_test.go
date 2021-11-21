package array

import (
	"testing"
)

// https://www.youtube.com/watch?v=v3WqNLmmBdk
// 思考漏的一个点是，
// 什么才算 top ?
//


// 1、思维点： 反转数组
// 2、


func minCostClimbingStairs(cost []int) int {
	f1 := 0
	f2 := 0
	f0 := 0
	for i := len(cost)-1;i >= 0;{
		f0 = cost[i] + min(f1, f2)
		f2 = f1
		f1 = f0
		i--
	}
	return min(f1, f2)
}

//func min(a, b int) int {
//	if a > b {
//		return b
//	} else {
//		return a
//	}
//}

func TestLen(t *testing.T){
	t.Log(len("https://github.com/sundayfun/daycam-server/pull/2461https://github.com/sundayfun/daycam-server/pull/2461https://github.com/sundayfun/daycam-server/pull/2461"))
}