package test

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T){
	ss := removeDuplicates("deeedbbcccbdaa", 3)
	t.Log(ss)
}

// AC 的答案，耶耶耶✌️
func removeDuplicates(s string, k int) string {
	// 递归删除
	// 递归结束条件是什么？
	// 怎么处理这个 k ?
	// Adjacent Duplicates
	stack := make([]byte, 0, len(s))
	for i := range s {
		count := 1
		// 要有一个后退策略
		// slen := len(stack)

		// for count <= k {

		// edge case
		// [100 100]
		// 1
		// [100 100 100]
		// 1
		for len(stack) +1 >= k && count < k {
			if stack[len(stack)-count] == s[i] {
				count++
			} else {
				// 这个不写会死循环
				break
			}
			fmt.Println("c",count)
			// 应该是要提前 break 一下
		}
		fmt.Println(stack)
		fmt.Println(count)
		// "deeedbbcccbdaa"
		// 不 ac
		if count == k {
			stack = stack[:len(stack)-count+1]
			// fmt.Println(string(stack))
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}