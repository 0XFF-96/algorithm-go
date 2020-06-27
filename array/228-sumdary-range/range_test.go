package main

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	// sorted
	// without duplicates

	// 用 map 来？
	var ret []string
	// [0, 3, 4, 5]
	// [1, 2, 3, 6]
	// [1, 2, 3, 6, 7]
	// [1, 2, 3, 5, 9, 10]
	// [1, 2, 3, 5, 9, 10, 11] 相关情况。
	// 有上面的三种情况。
	// edge case 处理
	//

	m := map[int]int{}


	// 哨兵节点？
	//

	previous := nums[0]
	var count int
	for i:=0; i < len(nums); i ++{
		if nums[i] != previous && nums[i] -1 !=  previous {
			m[previous] = count
			count = 0
		}
		count++
	}

	// for 循环一遍即可以
	//

	// previous := nums[0]
	// count := 0
	// for i:=1; i < len(nums); {
	//     if nums[i] - 1 != previous {
	//         if count == 0 {
	//             ret = append(ret, fmt.Printf("%s",strings.Iota(previous)))
	//         } else {
	//             ret = append(ret, fmt.Printf("%s->%s", strings.Iota(previous), nums[i]))
	//         }
	//         i += 2
	//         previous  = nums[i]
	//     }
	// }
	return ret
}


// AC 的答案
// 相关情况
func summaryRangesV2(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	var res []string
	f := nums[0]

	// len(nums) -1
	// 这里是解决边界条件很关键的一点🕐
	//
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] != nums[i]+1 {
			res = appendRange(res, f, nums[i])
			f = nums[i+1]
		}
	}
	return appendRange(res,f, nums[len(nums)-1])
}

func appendRange(str []string, f, n int) []string {
	s := strconv.Itoa(f)
	if f != n {
		s += "->" + strconv.Itoa(n)
	}
	return append(str, s)
}