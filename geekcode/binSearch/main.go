package binSearch


// 二分查找
// 1、循环退出条件 2、mid 的取值 3、low 和 high 的更新
func TestBinary(t *testing.T){
	a := []int{1, 2, 3, 4, 5}
	index := binarySearch(a, 4)
	t.Log(index)
}

// return the index
func binarySearch(sortedArray[]int,target int) int {
	low := 0
	high := len(sortedArray) -1

	for ;low < high;{
		mid := low + (high-low)>>1
		if sortedArray[mid] == target {
			return mid
		} else if sortedArray[mid] < target{
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func rBiSearch(sortedArray[]int, target ,low, high int) int {
	if (low > high) { return -1}
	mid := low + (high - low) >> 1
	if sortedArray[mid] == target {
		return mid
	} else if sortedArray[mid] < target{
		return rBiSearch(sortedArray, mid+1, high, target)
	} else {
		return rBiSearch(sortedArray, low, mid-1, target)
	}
}

// 二分搜索的相关条件
// 1、是否存在重复元素
// 2、查找第一个值等于给定值的元素。
// 当 a[mid] == value 时，大于，小于，等于 三种关系。 不能直接返回，而时需要继续搜索。
// 123444456 ,构造不同的测试用例。测试边界条件
// 停止条件是什么？


// 3、查找最后一个值等于给定值的元素
// 4、查找最后一个小于等于给定值的元素

// 5、如果查找的数据是一个循环的有序数组？
// 4， 5， 6， 1， 2， 3
//

// 查找第一个出现的元素的索引
func bSearch(a []int, n int, value int) int {
	low := 0
	high := n -1
	for ; low < high;{
		mid := low + (high-low) >> 1
		if a[mid] >= value {
			high = mid -1
		} else {
			low = mid + 1
		}
	}
	if low < n && a[low]==value {
		return low
	} else {
		return -1
	}
}

// 查找第一个出现的元素的索引
func bSearch2(a []int, n int, value int)int {
	low := 0
	high := n -1
	for ;low <= high; {
		mid := low + (high - low) >>1
		if a[mid] > value {
			high = mid -1
		} else if a[mid] < value {
			low = mid + 1
		} else {
			if mid == 0 || a[mid-1] != value {
				return mid
			} else {
				high = mid -1
			}
		}
	}
	return -1
}
// 查找最后一个值等于给定值的元素
// 查找第一个大雨等于给定值的元素
//

// skip list 相关的数据结构
// work flow
// 应用场景
// 1、作为一个种动态数据结构，我们需要用某种手段来维护索引与
// 原始链表大小之间的平衡。
// 2、重点在于：当不停地往跳表中插入数据的时候，
// 如果我们不更新索引，就有可能出现 2 个索引节点之间的数据非常多的情况
// 3、当我们往跳表中插入数据的时候，我们可以选择同时往这个
