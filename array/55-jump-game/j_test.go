package main

import (
	"testing"
)

func TestJumpGame(t *testing.T){
	res := canJumpV2([]int{2,3,1,1,4})
	t.Log(res)
}


// ac 从后往前的答案
// 
func canJump(nums []int) bool {
    lastPos := len(nums) -1 
    
    for i:= len(nums)-1; i >=0; i--{
        if i + nums[i] >= lastPos {
            lastPos = i
        }
    }
   return lastPos == 0
}
