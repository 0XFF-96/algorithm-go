package array

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	// find all
	// unique
	// dfs 的题目

	// cond: 可重复使用元素
	// limitation: 返回 set
	// 搜索空间 shrink


	// 问题，
	// 什么条件，终止了
	// 什么条件下，还能继续搜索？
	// target - cur > pick
	// 怎么才算是遍历完了这个数字
	//

	// 想到一个方法
	// cur = anyone in candidates []int
	// diff := target - cur
	// find out if there is any number >= diff
	// if true: then we pick the number,
	// 1.1 if the number equal to the tareget sum, then we put the number to our solution set.
	// else we move to the next number. which mean it is impossible to


	// sort.Int() O(nlogn)
	// pick number
	// see diff
	// [1, 2, 3, 6, 7]
	var ret [][]int
	sort.Ints(candidates)
	for _, c := range candidates {
		_ = c
		//s,ok := combinationSumH(candidates, tareget, cur)
		//if ok {
		//	ret = append(ret, s)
		//}
	}
	return ret
}

// 排列需要用 set 去重
// 组合并不需要 set 去重
// 利用【组合】的模板来实现
// 1、什么时候结束 2、如何实现能取重复元素
//
// https://leetcode.com/problems/combination-sum/discuss/432145/Go-golang-0ms-four-solutions
// 四种算法
//

// 这里有启发的意义
// https://leetcode.com/problems/combination-sum/discuss/455588/Golang-DFS-solution
func combinationSumV1(candidates []int, target int) [][]int {
	var ans [][]int
	var cur []int
	sort.Ints(candidates)
	// dfs(candidates, target, 0, &cur, &ans)

	var dfs func(candidates []int, target int, s int)

	dfs = func(candidates []int, target int, s int){
		if target == 0 {
			ans = append(ans, cur)
			return
		}

		for i:=s; i < len(candidates); i++{
			if candidates[i] > target {
				break
			}
			cur = append(cur, candidates[i])

			// i 尤其重要
			dfs(candidates, target - candidates[i], i)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(candidates, target, 0)
	return ans
}

// 看看怎么 combination 怎么做。
// 各种变形。
//