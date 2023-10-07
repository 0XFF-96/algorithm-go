package array

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(index))
	for i := range res {
		res[i] = -1
	}

	for i, indexI := range index {
		if res[indexI] != -1 {
			copy(res[indexI+1:], res[indexI:])
		}
		res[indexI] = nums[i]
	}
	return res
}

// 很蠢的用复制特性...
// 最好的方法，应用是用链表。然后，链表转数组。
// 利用数组实现链表特性
// 考虑这两种情况
// 1, 2, 3, 4, 5
// 5, 4, 3, 2, 1,
// 交换排序
// 移动位置。
