
func findDuplicate(nums []int) int {
    tmp := make([]int, len(nums))
    for _, v := range nums {
        tmp[v]++
        if tmp[v] > 1 { return v }
    }
    return -1
}
