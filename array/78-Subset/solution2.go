package _8_Subset


// 递归的方式
//
func subsetsV1(nums []int) [][]int {
subsets := [][]int{[]int{}}
for _, num := range nums {
for _, subset := range subsets {
newSubset := append([]int{}, subset...)
newSubset = append(newSubset, num)
subsets = append(subsets, newSubset)
}
}
return subsets
}
