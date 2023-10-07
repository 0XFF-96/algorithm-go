package reorderarray

import (
	"testing"
)

// 排序的思想
// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
// 使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，
// 并保证奇数和奇数，偶数和偶数之间的相对位置不变*
func TestReorderArray(t *testing.T){
	array := []int{1, 2, 3, 4, 5, 6, 7}
	n := reorderArray(array)
	t.Log(n)
}

// 不行
// 这个顺序乱了
// 只能开辟数组空间来保证
// ?
func reorderArray(array []int)[]int{
	if len(array) == 0{
		return nil
	}
	lright := len(array) -1
	// lright := l
	for i:=0; i <lright;{
		if array[i] %2 == 0 {
			array[lright], array[i] = array[i], array[lright]
			lright--
			continue
		}
		i++
	}
	return array
}

