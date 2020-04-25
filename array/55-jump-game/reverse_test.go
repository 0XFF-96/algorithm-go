// 逆向思维
// 这种结果很优美啊

// Runtime: 8 ms, faster than 91.19% of Go online submissions for Jump Game.
// Memory Usage: 4.2 MB, less than 85.71% of Go online submissions for Jump Game.
func canJump(nums []int) bool {
    lastPos := len(nums) -1 
    
    for i:= len(nums)-1; i >=0; i--{
        if i + num[i] >= lastPos {
            lastPos = i
        }
    }
   return lastPos == 0
}