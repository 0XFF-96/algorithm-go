package array

import (
	"fmt"
	"sort"
	"testing"
)


// https://leetcode.com/problems/combination-sum-ii/

// 在for循环里剪枝（return），和 递归时候 当成base case return 有什么区别吗
//

// Set是偷懒的做法
// 9:14 讲的第二种方法从本质上去除重复

// how to remove duplicates ?
// - Use set ? golang 中怎么用？有很多个数组，数字唯一性
// - Disallow same number in the same depth.
//
func TestCombinationSumV2(t *testing.T){
	res := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	t.Log(res)
	// 产生重复是因为，candidates 有重复元素
	//
}

// https://www.youtube.com/watch?v=RSatA4uVBDQ&t=2s
// 不 AC, 差🕐
// 有一个部分内容,
// 1-做到能不使用重复元素
// 2-做不到 set 去重
// backtracking 的时候，没有做去重操作...


// 1、一开始写了一个错的条件， i 没有加 1
// 2、pos 条件填错
// 3、去除重复元素的条件写错。
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	var res [][]int
	var cur []int

	// set := map[int]struct{}{}
	sort.Ints(candidates)
	fmt.Println("can", candidates)
	var dfs func(candidates []int, target int, pos int)

	dfs = func(candidates []int, target int, pos int) {
		if target == 0 {
			fmt.Println("----------")
			fmt.Println("res", res)
			fmt.Println("cur", cur)

			//s := sum(cur)
			//if _, ok := set[s]; ok {
			//	return
			//}
			res = append(res, append([]int{}, cur...))
			// set[s] = struct{}{}
			fmt.Println("----------")
			return
		}

		for i := pos; i < len(candidates); i++{
			if candidates[i] > target {
				break
			}

			//if i > 0 && candidates[i] == candidates[i-1] {
			//	continue
			//}

			if i > pos && candidates[i] == candidates[i-1] {
				continue
			}
			cur = append(cur, candidates[i])
			dfs(candidates, target - candidates[i], i+1)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(candidates, target, 0)
	return res
}

// 这个也是错的？
//
func combinationSum2v1(candidates []int, target int) [][]int {
	var ans [][]int
	var cur []int
	sort.Ints(candidates)

	var dfs func(candidates []int, target int, s int)

	dfs = func(candidates []int, target int, s int){
		if target == 0 {
			ans = append(ans, append([]int{}, cur...))
			return
		}

		for i:=s; i < len(candidates); i++{
			if candidates[i] > target {
				break
			}

			// 这语句的排列次数不对啊大哥！
			//
			cur = append(cur, candidates[i])

			// i 尤其重要
			// 为什么不是 i+1?

			if i > s && candidates[i] == candidates[i-1] {
				continue
			}
			dfs(candidates, target - candidates[i], i+1)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(candidates, target, 0)
	return ans
}

// 这是正确答案，
// 对比上面，做对了哪一步？
func combinationSum2v(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	var res [][]int
	var cur []int

	sort.Ints(candidates)
	var dfs func(candidates []int, target int, pos int)

	dfs = func(candidates []int, target int, pos int) {
		if target == 0 {
			res = append(res, append([]int{}, cur...))
			return
		}

		for i := pos; i < len(candidates); i++{
			if candidates[i] > target {
				break
			}

			if i > pos && candidates[i] == candidates[i-1] {
				continue
			}
			cur = append(cur, candidates[i])
			dfs(candidates, target - candidates[i], i+1)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(candidates, target, 0)
	return res
}


func sum(a []int)int {
	var res int
	for _, i := range a {
		res += i
	}
	return res
}


// 为什么会失败？
//func combinationSum2(nums []int, target int) (result [][]int) {
//	sort.Ints(nums)
//	combinationSum2Helper(nums, nil, target, 0, 0, &result)
//	return result
//}
//
//func combinationSum2Helper(nums, combo []int, target, sum, startIndex int, result *[][]int) {
//	if sum == target {
//		*result = append(*result, append([]int{}, combo...))
//		return
//	}
//	for i := startIndex; i < len(nums) && (sum + nums[i]) <= target; i++ {
//		if i != startIndex && nums[i] == nums[i - 1] { continue }
//		combinationSum2Helper(nums, append(combo, nums[i]), target, sum + nums[i], i + 1, result)
//	}
//}

