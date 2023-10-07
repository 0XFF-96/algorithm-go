package main

import (
	"fmt"
)

type ListNode struct {
	     Val int
	     Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	// 需不需要辅助空间？
	array := []*ListNode{}
	cur := head
	for cur != nil {
		fmt.Println("cur.Val", cur.Val)
		array = append(array, &ListNode{Val: cur.Val})
		cur = cur.Next
	}
	fmt.Println("a1.val", array[2].Val)
	// inplement insertion sort
	// insertion sort 模板？
	// inserttion sort 的坑和优化点？
	// 重点
	for i:=0; i < len(array); i++{
		// unsort
		unsort := array[i]
		fmt.Println("unsort", array[i].Val)
		for j := i; j > 0; j -- {
			if unsort.Val < array[j-1].Val {
				array[j] = array[j-1]
			} else {
				array[j] = unsort
				break
			}
		}
		fmt.Println("unsort", array[i].Val)
	}
	// 上面部分，代码逻辑有问题。
	// 看看是什么问题？
	//

	// reformat the array
	dummy := &ListNode{}
	prev := dummy
	fmt.Println("len", len(array))
	for i :=0; i < len(array); i++ {
		fmt.Println("a.Val", array[i].Val)
		prev.Next = array[i]
		prev = prev.Next
	}

	return dummy.Next
}

func InsertionSort(nums []int) []int {
	n := len(nums)
	// Run the outer loop n - 1 times, from index 1 to n-1, as first element is already sorted
	// At the end of ith iteration, we have sorted list [0, i]
	for i := 1; i <= n-1; i++ {
		// Pick ith element and keep swapping with i-1th element if nums[i] < nums[i-1]
		j := i
		for j > 0 {
			// If value at index j is smaller than the one at j-1, swap them
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
			j -= 1
		}
	}
	return nums
}

func insertionSortListV2(head *ListNode) *ListNode {
	// 需不需要辅助空间？
	array := []*ListNode{}
	cur := head
	for cur != nil {
		array = append(array, cur)
		cur = cur.Next
	}
	// inplement insertion sort
	// insertion sort 模板？
	// inserttion sort 的坑和优化点？
	// 重点
	for i:=0; i < len(array); i ++{
		// unsort
		j := i
		for j > 0 {
			// 这里的逻辑!!
			if array[j].Val < array[j-1].Val {
				fmt.Println(array[j].Val)
				fmt.Println(array[j-1].Val)
				array[j].Val, array[j-1].Val = array[j-1].Val, array[j].Val
			}
			j -= 1
		}
	}

	// reformat the array
	dummy := &ListNode{}
	prev := dummy
	for i :=0; i < len(array); i++ {
		prev.Next = array[i]
		prev = prev.Next
	}

	return dummy.Next
}