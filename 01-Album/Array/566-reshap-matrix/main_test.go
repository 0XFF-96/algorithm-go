

// 
func matrixReshape(nums [][]int, r int, c int) [][]int {
    if len(nums) == 0 || r*c != len(nums) * len(nums[0]) {
        return nums 
    }
    ret := make([][]int, r)
    for i:=0;i<r;i++{
        ret[i] = make([]int, c)
    }
    
    count := 0 
    for i:=0 ; i <len(nums); i++{
        for j := 0;j < len(nums[0]);j++{
            ret[count/c][count % c] = nums[i][j]
            count++
        }
    }
    return ret 
}
