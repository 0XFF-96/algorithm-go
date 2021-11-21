package array

import (
	"testing"
)

// 1, 0, 0, 0
// 0, 0, 0, 1
// 这两种极端情况


// 1、循环条件写错
// 2、返回条件，没有返回它想要的条件。
//
func TestMaxDist(t *testing.T){

	a := maxDistToClosest([]int{1,0,0,0,1,0,1})
	t.Log(a)
}

func maxDistToClosest(seats []int) int {
	if len(seats) == 0 {
		return 0
	}
	// need to pass
	// Left to Right
	// Right to Left
	LR := make([]int, len(seats))
	RL := make([]int, len(seats))

	leftMostIdx := -1000000000
	rightMostIdx := 1000000000

	for i:=0;i<len(seats);i++{
		if seats[i] == 1 {
			leftMostIdx = i
		} else {
			LR[i] = i - leftMostIdx
		}
	}

	for i:= len(seats)-1;i>=0;i--{
		if seats[i] == 1 {
			rightMostIdx = i
		} else {
			RL[i] = rightMostIdx - i
		}
	}

	maxDistance :=0
	// idx := 0
	for i:=0; i < len(seats);i++{
		mmin := min(LR[i], RL[i])
		if mmin > maxDistance {
			maxDistance = mmin
			// idx = i
		}
	}
	return maxDistance
}


//func min(a, b int)int{
//	if a > b {
//		return b
//	} else {
//		return a
//	}
//}


//
//func min(a, b int)int{
//	if a > b {
//		return b
//	} else {
//		return a
//	}
//}



