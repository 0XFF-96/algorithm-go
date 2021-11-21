package find_number_from_2d_array

import (
	"testing"
)

// 二维数组查找
// https://www.nowcoder.com/questionTerminal/abc3fe2ce8e146608e868a70efebf62e?answerType=1&f=discussion
// https://www.nowcoder.com/questionTerminal/abc3fe2ce8e146608e868a70efebf62e?answerType=1&f=discussion
// 矩阵是有序的
// 可以利用的信息：矩阵是有序的。
// 利用二维数组由上到下，由左到右递增的规律
// 将每一行看作递增数组，然后使用二分搜索
// 时间复杂读 O(行宽+列高）
func TestFindNumFrom2dArray(t *testing.T){
	array := [][]int{
		{
			1,2,3,4,5,
		},
		{
			5,6,7,8,9,
		},
		{
			10, 11, 12,14,15,
		},
	}
	// 需要一定机制保证输入的 array 必须为数组。不能出现空列的情况
	res := bruteForceTravel2DArray(10, array)
	res2 := searchFromRightCorner(10, array)
	res3 := searchFromRightCorner2(10, array)
	t.Log(res)
	t.Log(res2)
	t.Log(res3)
}

func bruteForceTravel2DArray(target int, array [][]int)bool{
	// 遍历二维数组的写法
	if len(array) == 0 {
		return false
	}
	for idx := range array {
		for _, val := range array[idx]{
			if val == target {
				return true
			}
		}
	}
	return false
}

func searchFromRightCorner(target int, array[][]int)bool{
	row := len(array)
	if row == 0 {
		return false
	}
	// 检查列是否为空
	cols := len(array[0])
	if cols == 0 {
		return false
	}

	row = row -1
	col := 0

	//fmt.Println(row)
	//fmt.Println(cols)
	for ;row >=0 && col < cols;{
		if(array[row][col] < target) {
			col++;
		} else if array[row][col] > target {
			row--;
		} else {
			return true
		}
	}
	return false
}
