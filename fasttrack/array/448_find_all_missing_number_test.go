package array


//
// 如何利用下标信息 mark negetive
// 利用位运算知识？
//

// 设置某一位为 1
// 找到某一位，不为 1 的元素。
//

// 相当于位运算的做法
func findDisappearedNumbers(nums []int) []int {
	length := len(nums)

	n := make([]int, length)

	for _, v := range nums {
		n[v-1] = -v
	}
	var ret []int
	for idx, _ := range nums {
		if n[idx] == 0 {
			ret = append(ret, idx+1)
		}
	}
	return ret
}
