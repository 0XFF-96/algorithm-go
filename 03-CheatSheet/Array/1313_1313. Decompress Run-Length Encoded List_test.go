package array


func decompressRLElist(nums []int) []int {
    if len(nums) == 0 || len(nums) % 2 != 0{
        return nil 
    }
    var ret []int
    for i:=0;i < len(nums);{
        freq := nums[i]
        val := nums[i+1]
        for j:=0;j<freq;j++{
            ret = append(ret, val)
        }
        i +=2 
    }
    return ret
}