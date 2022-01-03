package main

import (
	"sort"
	"testing"
)

func TestCombinationSum1(t *testing.T){
	res := combinationSum([]int{2, 3, 6,4, 7}, 7)
	t.Log(res)

	res = combinationSumV2([]int{2, 3, 6, 7}, 7)
	t.Log(res)
}



// 允许使用重复元素
//
func combinationSum(candidates []int, target int) [][]int {
	r := [][]int{}
	for i, c := range candidates {
		if c == target {
			r = append(r, []int{c})
		} else if c < target {
			//
			// candidates[i+1:], target -c
			// 还可以这样写
			// 这样就不会允许重复元素出现
			for _, rr := range combinationSum(candidates[i:], target - c) {
				r = append(r, append(rr, c))
			}
		}
	}
	return r
}

//candidates = [2,3,6,7], target = 7,
//[
//[7],
//[2,2,3]
//]
// limitation: The solution set must not contain duplicate combinations
// https://leetcode.com/problems/combination-sum/
func combinationSumV2(candidates []int, target int) [][]int {
	var ans [][]int
	var cur []int
	// 为什么需要排序？
	//
	sort.Ints(candidates)
	// dfs(candidates, target, 0, &cur, &ans)

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
			cur = append(cur, candidates[i])

			// i 尤其重要
			// 为什么不是 i+1?
			// 因为可以使用重复元素
			dfs(candidates, target - candidates[i], i)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(candidates, target, 0)
	return ans
}


func TestConbinationSum3(t *testing.T){
	//ret := combinationSum3(3, 7)
	//t.Log(ret)


	ret2 := combinationSum3V2(3, 7)
	t.Log(ret2)
}

// 如果是使用匿名函数，
// 会追踪不到，栈堆的相关信息💻。
//

// 成功了
// AC 了
// 根据花花的答案写的。
// 为什么 dfs 执行了 19 次？
// 你能否白板，画出调用执行流程图？
//
func combinationSum3(k int, n int) [][]int {
	var ret [][]int

	// 辅助理解这个函数在干什么的参数
	var count int
	var dfs func(k int, n int, s int, curSum []int)

	dfs = func(k int ,n int, s int, cur []int) {
		if k == 0 {
			if n == 0 {
				// 这个操作有点危险⚠️
				// ret = append(ret, cur)
				ret = append(ret, append([]int{}, cur...))
			}
		}
		// fmt.Println("------cur------")
		// fmt.Println(cur)
		count++
		// fmt.Println(count)
		for i := s ; i <= 9; i++{
			if i > n {
				return
			}
			cur = append(cur, i)

			// k 是什么，组合的深度
			// n - i, 结果
			// i + 1, 下一个可以利用的元素
			// cur 数组
			dfs(k - 1, n - i, i + 1, cur)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(k, n, 1, []int{})
	return ret
}

func combinationSum3V2(k int, n int) [][]int {
	var ret [][]int
	helper3v2(k, n, 1, []int{}, &ret)
	return ret
}

func helper3v2(k int, n int, s int, cur []int, ret *[][]int){
	if k == 0 {
		if n == 0 {
			// 这个操作有点危险⚠️
			// ret = append(ret, cur)
			*ret = append(*ret, append([]int{}, cur...))
		}
	}
	// fmt.Println("------cur------")
	// fmt.Println(cur)
	// count++
	// fmt.Println(count)
	for i := s ; i <= 9; i++{
		if i > n {
			return
		}
		cur = append(cur, i)

		// k 是什么，组合的深度
		// n - i, 结果
		// i + 1, 下一个可以利用的元素
		// cur 数组
		helper3v2(k - 1, n - i, i + 1, cur, ret)
		cur = cur[:len(cur)-1]
	}
}



// 有可能 go 没办法追踪匿名函数的调用
//


