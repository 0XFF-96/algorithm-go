package findMinNumberInRotateArray

//func TestSpinArray(t *testing.T){
//
//	testCase := []struct{
//		input []int
//		output int
//	}{
//		{
//			input:[]int{1, 2, 3, 4, 5},
//			output:1,
//		},
//		{
//			input:[]int{3, 4, 5, 1, 2},
//			output:1,
//		},
//		{
//			input:[]int{3, 3, 3, 3, 1, 3},
//			output:1,
//		},
//		{
//			input:[]int{1, 1, 1, 1, 0, 1},
//			output:0,
//		},
//	}
//	for _, tt := range testCase{
//		res := findMinNumberInRotateArray(tt.input)
//		require.Equal(t, res, tt.output)
//		// t.Log(res)
//	}
//	// 修复 index out of range 问题
//}

func findMinNumberInRotateArray(n []int) int {
	lLen := 0
	rLen := len(n)-1
	for ;lLen < rLen;{
		if n[lLen] < n[rLen] {
			// 没有被旋转的时候
			return n[lLen]
		}
		// 二分搜索的变形，决定向哪一个方向收缩
		mid := (lLen + rLen) >> 1
		if n[mid] > n[lLen] {
			lLen = mid + 1
		} else if n[mid] < n[rLen] {
			// rLen = mid - 1 ? 会如何？
			rLen = mid
		} else {
			lLen++
		}
	}
	return n[lLen]
}


//public class Solution {
//public int minNumberInRotateArray(int[] array) {
//int i = 0, j = array.length - 1;
//while (i < j) {
//if (array[i] < array[j]) {
//return array[i];
//}
//int mid = (i + j) >> 1;
//if (array[mid] > array[i]) {
//i = mid + 1;
//} else if (array[mid] < array[j]) {
//j = mid; // 如果是mid-1，则可能会错过最小值，因为找的就是最小值
//} else i++;  // 巧妙避免了offer书上说的坑点（1 0 1 1 1）
//}
//return array[i];
//}
//}