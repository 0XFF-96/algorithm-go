package sort

import (
	"sort"
	"testing"
)

func reverse(s []int){
	for i,j := 0,len(s)-1;i < j;i,j = i+1,j-1{
		s[i], s[j] = s[j],s[i]
	}
}

func TestReverse(t *testing.T){
	// 区分这行的区别，分别是 数组和切片...
	//// a := [...]int{0, 1, 2, 3, 4, 5, 6}
	//a := []int{1,2,3,4,5,6}
	//reverse(a)
	//t.Log(a)
	t.Log("dummy code")
	b := [...]int{1,2,4}
	c := b[:]
	t.Log(c)
	// 变成一个切片，而对切的操作，是会影响底层数组的...
	reverse(b[:])
	capB := cap(b)
	lenB := len(b)
	t.Log(capB)
	t.Log(lenB)
	t.Log(b)
	s := []int{}
	reverse(s)
	t.Log(s)
	// 在调试看不到有函数的实现，都是在
	// 运行时，给相关函数加料的...
	//
}

func appendInt(x []int, y int)[]int{
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by dobbling, for amoritize linear complexity
		zcap := zlen
		if zcap < 2*len(x){
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z,x)
	}
	z[len(x)] = y
	return z
}

func TestAppendInt(t *testing.T){
	var x, y []int
	for i := 0;i < 10;i++{
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t\t%v\n", i, cap(y), y)
		x = y
	}
}

func TestSlice(t *testing.T){
	// 如果没有超过 slice 的 cap 是不会 panic 的
	ss := make([]int,2,10)
	t.Log(ss)
	// 下面这个函数会 panic
	// t.Log(ss[3])
	t.Log(ss[:9])
	// 区别在于，下面这个函数不会 panic, 因为底层数组没有超出...
	// t.Log(ss[:11])
}
// ...

func remove(slice []int, i int)[]int{
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int,i int)[]int{
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func TestSort(t *testing.T){
	unsort := []int{1,3,4,56,7,9654,2}
	b1 := bubbleSort(unsort)
	b2 := insertionSort(unsort)
	t.Log(b1)
	t.Log(b2)
}


// 冒泡排序
func bubbleSort(list []int)[]int{
	if len(list) <= 1 {
		return list
	}
	for i := range list {
		flag := false
		if list[i] > list[i+1] { // swap
			list[i], list[i+1] = list[i+1], list[i]
			flag = true
		}
		// 没有数据交换，提前退出
		if !flag {
			break
		}
	}
	return list
}

// 这样会造成数据的复制
// 输入 []int 和 []int
func insertionSort(list []int) []int {
	if len(list) <= 1{
		return list
	}
	for i := range list{
		tmp := list[i]

		// find the insert index
		j := i-1
		for ;j>=0; j--{
			if list[j] > tmp {
				list[j+1] = list[j]
			} else {
				break
			}
		}
		list[j+1] = tmp
	}
	return list
}

// merge_sort
// 相关例子

func merge_sort(a []int, l int) []int{

	return nil
}


// 这种性能应该很低，但是
func merge_sort_helper(a[]int,begin, end int) []int {
	if begin >= end {
		return a
	}
	mid := (begin+end) / 2
	a1 := merge_sort_helper(a,begin, mid)
	a2 := merge_sort_helper(a,mid+1, end)
	return append(a1, a2...)
}

// 利用哨兵，简化 merge 的代码...
// 1.稳定的排序算法 2.时间复杂度是多少 3.空间复杂度是多少
// 归并不是原地排序 o(n)
//

// 快排..
// 重点在与 找 pivot 点.
// ? 类似与选择排序？区分 未处理空间与已处理空间？
// 把排序的和未排序的数据，区分为两个 swap
// 每次，已 a[end] 为轴点...
func partition(a []int, begin, end int){
	pivot := a[end]
	i := begin
	for j := begin; j < end-1;j++{
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i = i + 1
		}
	}
	a[i], a[end] = a[end],a[i]
}

// 线性排序，它是基于非比较的排序算法。

// 桶排序的相关方法,1、分布到几个有序桶 2、桶之间数据不需要再排序
// 3、如果不能一次性读入内存，继续划分小区间


// 计数排序是桶排的一种特殊情况，它省略了桶内排序的时间
//

// 桶排序, 将 xxx 分成不同的桶，然后按照桶进行排序。
// 这个思路比较像...
// 2.1) 桶个数的设置依据什么原则?
// 2.2) 桶大小的设置，让桶的能动扩容?
// ？
//

func TestStandSort(t *testing.T){
	// 探究一下标准库中的排序例子
	a := []string{"4","1", "2", "b", "A", "a"}
	sort.Strings(a)
	t.Log(a)
}

// 插入排序

// Insertion sort
//func insertionSort(data Interface, a, b int) {
// 这里为什么是 a + 1 ?, 有可能和 j -- 有关系。
//	for i := a + 1; i < b; i++ {
//		for j := i; j > a && data.Less(j, j-1); j-- {
//			data.Swap(j, j-1)
//		}
//	}
//}


// sfit down 堆排堆一个算法

// siftDown implements the heap property on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
//func siftDown(data Interface, lo, hi, first int) {
//	root := lo
//	for {
//		child := 2*root + 1
//		if child >= hi {
//			break
//		}
//		if child+1 < hi && data.Less(first+child, first+child+1) {
//			child++
//		}
//		if !data.Less(first+root, first+child) {
//			return
//		}
//		data.Swap(first+root, first+child)
//		root = child
//	}
//}

// heap sort 堆排序
//func heapSort(data Interface, a, b int) {
//	first := a
//	lo := 0
//	hi := b - a
//
//	// Build heap with greatest element at top.
//	for i := (hi - 1) / 2; i >= 0; i-- {
//		siftDown(data, i, hi, first)
//	}
//
//	// Pop elements, largest first, into end of data.
//	for i := hi - 1; i >= 0; i-- {
//		data.Swap(first, first+i)
//		siftDown(data, lo, i, first)
//	}
//}

// 可以模仿官方库的例子，写相关算法

// 快排
//func quickSort(data Interface, a, b, maxDepth int) {
//	for b-a > 12 { // Use ShellSort for slices <= 12 elements
//		if maxDepth == 0 {
//			heapSort(data, a, b)
//			return
//		}
//		maxDepth--
//		mlo, mhi := doPivot(data, a, b)
//		// Avoiding recursion on the larger subproblem guarantees
//		// a stack depth of at most lg(b-a).
//		if mlo-a < b-mhi {
//			quickSort(data, a, mlo, maxDepth)
//			a = mhi // i.e., quickSort(data, mhi, b)
//		} else {
//			quickSort(data, mhi, b, maxDepth)
//			b = mlo // i.e., quickSort(data, a, mlo)
//		}
//	}
//	if b-a > 1 {
//		// Do ShellSort pass with gap 6
//		// It could be written in this simplified form cause b-a <= 12
//		for i := a + 6; i < b; i++ {
//			if data.Less(i, i-6) {
//				data.Swap(i, i-6)
//			}
//		}
//		insertionSort(data, a, b)
//	}
//}


func topoSort(m map[string][]string)[]string{
	var order []string
	seem := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string){
		for _, item := range items {
			if !seem[item]{
				seem[item] = true
				visitAll(m[item])  // 遍历其他的, 最重要是这一行递归，一定要放在 append(order,item)之前
				order = append(order,item)   // ? 为什么？
			}
		}
	}
	// 真正的实现
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func TestT(t *testing.T){
	// ticker := time.NewTicker(1)
	// ticker.C
	// websocket.PingMessage
}
