package array

// 有一个 edge case 没有过
// 当是一个 zero value 的时候，但是，有不能把它 copy
// https://leetcode.com/problems/duplicate-zeros/solution/
//
// in-place 需要移动数组元素等特征的题目时，如果我们从前往后
// 进行元素复制，往往有坑。这时候我们可以考虑从后往前复制。
// 这招屡试不爽
func duplicateZeros(arr []int)  {
	var start int
	count := 0

	// 在这里没有处理，本来就应该去掉的情况
	// 如果有这种情况，
	// 不应该从最后一个元素开始
	// 而是，最后一个放 0
	// 从最后一个 0 元素开始遍历。
	for idx, v := range arr {
		if v == 0 {
			count += 2
		}
		count ++
		if count >= len(arr) -1 {
			start = idx
			break
		}
	}
	// need to note where it stops
	// 1 2 3 4 5 0 0
	// 0 0 1 2 3 4 5
	// 0 1 0 2 3 4 5
	// [8,4,5,0,0,0,0,7]
	// 多了一个 0 ...
	// edge case
	l := len(arr)-1
	// fmt.Println(start)
	for i:= start; i >=0;i--{
		if arr[i] == 0 {
			arr[l] = 0
			arr[l-1] = 0
			l -= 2
			continue
		}
		arr[l] = arr[i]
		l--
	}
}

// 看看其他答案？
// https://leetcode.com/problems/duplicate-zeros/discuss/312781/Golang-Space-O(1)-RecursiveOne-pass-and-Space-O(N)
//
//
// 需要注意⚠️的情况：Handle the edge case for a zero present on the boundary of the leftover elements
//



// 吊....
func duplicateZerosV2(arr []int)  {
	for i:= 0; i < len(arr); i++ {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:])
			i++
		}
	}
}

// 效率很低..
// O(N) 效率
func duplicateZerosV3(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		if arr[i] == 0 {
			for j := length - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			i++
		}
	}
	return arr
}

