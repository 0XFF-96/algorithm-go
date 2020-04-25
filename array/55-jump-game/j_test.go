package main

import (
	"testing"
)

func TestJumpGame(t *testing.T){
	res := canJumpV2([]int{2,3,1,1,4})
	t.Log(res)
}
