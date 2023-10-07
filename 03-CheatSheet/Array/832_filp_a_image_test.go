package array

import (
	"testing"
)

func TestFilpAndInvertIamge(t *testing.T){
	res := flipAndInvertImage([][]int{{1,1,0},{1, 0, 1},{0, 0, 0}})
	t.Log(res)

	// golang 共享底层数组
}

func TestArray(t *testing.T){
	a := []int{1,2,3}
	b := a
	a[2] = 0
	t.Log(b)
}


func flipAndInvertImage(A [][]int) [][]int {
	// flip image horizontally
	//
	//

	// 找到一个规律
	// 如果开始为1，最后一定为 0
	// [1, 1, 0] -> [1,0,0]
	// 从原始数组开始，写入原始数组末尾

	// 数组遍历框架
	//
	// 数据在遍历的时候，被篡改了...
	//
	row := len(A)
	col := len(A[0])
	// 能不能写一个就地算法？
	// var tmp []int
	// tmp := []int{}
	for r:=0; r < row; r++{
		// 这一行 deep copy 出现了问题..
		tmp := append([]int{}, A[r]...)
		// copy(tmp, A[r])
		// fmt.Println(tmp)


		for c := 0; c < col; c++{
			if tmp[c] == 0 {
				A[r][col- c -1] = 1
			} else {
				A[r][col -c -1] = 0
			}
		}
	}
	return A
}

// 为什么我的答案会错？

//class Solution {
//public int[][] flipAndInvertImage(int[][] A) {
//int C = A[0].length;
//for (int[] row: A)
//for (int i = 0; i < (C + 1) / 2; ++i) {
//int tmp = row[i] ^ 1;
//row[i] = row[C - 1 - i] ^ 1;
//row[C - 1 - i] = tmp;
//}
//
//return A;
//}
//}



// 这里有个坑
// 关于数组的
func flipAndInvertImageV2(A [][]int) [][]int {
	// flip image horizontally
	//
	//

	// 找到一个规律
	// 如果开始为1，最后一定为 0
	// [1, 1, 0] -> [1,0,0]
	// 从原始数组开始，写入原始数组末尾

	// 数组遍历框架
	//

	row := len(A)
	col := len(A[0])
	// 能不能写一个就地算法？

	// must need a deep copy of it
	// try to be fast ....
	for r:=0; r < row; r++{
		tmp := append([]int{}, A[r]...)
		for c := 0; c < col; c++{
			if tmp[c] == 0 {
				A[r][col- c -1] = 1
			} else {
				A[r][col -c -1] = 0
			}
		}
	}
	return A
}