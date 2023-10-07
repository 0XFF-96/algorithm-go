package array

import (
	"fmt"
	"testing"
)


// 这个开始的代码好想不对。
// 思路不正确🙆
func isMonotonic(A []int) bool {
	if len(A) < 2 {
		return true
	}

	desc := false
	if A[0] > A[1] || A[len(A)-2] > A[len(A)-1]{
	    desc = true
	}

	for i :=1; i < len(A); i++{
		if A[i-1] < A[i] && desc {
			return false
		} else if A[i-1] > A[i] && !desc {
			return false
		}
	}
	return true
}

func TestPrint(t *testing.T){
	t.Log(fmt.Sprintf("ahahha"))
}

func isMonotonicV2(A []int) bool {
	if len(A) < 2 {
		return true
	}
	return increasing(A) || decreasing(A)
}

func increasing(A []int) bool {
	for i:=0;i < len(A)-1;i++{
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}

func decreasing(A []int) bool {
	for i:=0;i < len(A)-1;i++{
		if A[i] < A[i+1] {
			return false
		}
	}
	return true
}
