package array

import (
	"sort"
	"testing"
)

// 没有过 test case
// 双指针的情况下。 two pointers
// 需要好好总结这类题目的特点。


// https://leetcode.com/problems/squares-of-a-sorted-array/discuss/222079/Python-O(N)-10-lines-using-deque-beats-100
// 这里的相关代码
// two pointers
// pointers 在什么情况下，移动，怎么移动，什么时候结束？
//
func sortedSquares(A []int) []int {
	// 三种情况，three cases:
	// 全部 negetive
	// 全部 positive
	// 一半一半。hack。

	//     if A[0] >=0 {
	//         for idx, v := range A {
	//             A[idx] = square(v)
	//         }
	//         return A
	//     }

	//     if A[len(A)-1] <= 0 {
	//         for idx, v := range A {
	//             A[idx] = square(v)
	//         }
	//         return A
	//     }

	l := 0
	r := len(A) -1

	for l <= r {
		ls := square(A[l])
		rs := square(A[r])

		if ls >= rs {
			A[l] = A[r]
			A[r] = ls
			r--
		} else {
			A[r] = rs
			r--
		}
	}

	return A
}


func square(num int) int {
	return num * num
}


func sortedSquaresV2(A []int) []int {
	if len(A) == 0 {
		return nil
	}
	for idx, v := range A {
		A[idx] = v * v
	}

	sort.Ints(A)
	return A
}

func TestE(t *testing.T){
	num := exchange([]int{1, 2, 3, 4, 5})
	t.Log(num)
}

func exchange(A []int) []int {
	i := 0
	j := len(A) -1

	for i < j {
		A[i],A[j] = A[j],A[i]
		i++
		j--
	}
	return A
}

func sortedSquaresV3(A []int) []int {
	n := len(A) -1
	j := 0

	var ret []int
	for j <= n {
		left, right := absSigle(A[j]), absSigle(A[n])
		if left > right {
			ret = append(ret, left * left)
			j++
		} else {
			ret = append(ret, right * right)
			n--
		}
	}

	return reverseArray(ret)
}

func absSigle(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}


func reverseArray(A []int) []int {
	i := 0
	j := len(A) -1

	for i < j {
		A[i],A[j] = A[j],A[i]
		i++
		j--
	}
	return A
}

// 看一下这个逻辑。
//    while (l < r) {
//        int x = A[l] * A[l];
//        int y = A[r-1] * A[r-1];
//
//        if (x >= y) {
//            result[p] = x;
//            p--;
//            l++;
//        } else {
//            result[p] = y;
//            p--;
//            r--;
//        }
//    }