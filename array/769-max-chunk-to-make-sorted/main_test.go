package _69_max_chunk_to_make_sorted

import (
	"fmt"
	"testing"
)
// [1, 2, 0, 3]
// 这种情况？
// [2, 0, 1], 这种 case 已经考虑过了。


// 为什么我的代码会不过？
func TestMaxChunk(t *testing.T) {
	ret := maxChunksToSorted([]int{2, 0, 1})
	t.Log(ret)
}

func maxChunksToSorted(arr []int) int {
	// 1. 0 - n-1 范围内
	// 2.
	// 3. xxx
	// max 动态规划？
	//
	if len(arr) <= 1 {
		return len(arr)
	}

	// merge & split 两个操作。
	// merge: 前一个数字，比后一个数字大。 （指针维持不变）
	// split: 前一个数字，比后一个数字少。 （+1， 指针向后移动）
	// 两个极端例子。

	count := 1
	curIdx := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			// merge
		} else {
			// split
			if arr[curIdx] <= arr[i] {
				curIdx = i
				count++
				fmt.Println(arr[curIdx], arr[i])
			}
		}
	}
	// edge case
	return count
}

// AC
// 有一个点没有确定好，导致这个题有个问题。
//
func maxChunksToSortedV2(arr []int) int {
	// 1. 0 - n-1 范围内
	// 2.
	// 3. xxx
	// max 动态规划？
	//
	if len(arr) <= 1 {
		return len(arr)
	}

	// merge & split 两个操作。
	// merge: 前一个数字，比后一个数字大。 （指针维持不变）
	// split: 前一个数字，比后一个数字少。 （+1， 指针向后移动）
	// 两个极端例子。


	ans := 0
	m := 0
	for i := 0; i < len(arr); i++{
		m = max(m, arr[i]);
		if m == i {
			ans++
		}
	}
	return ans
}

func max(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}