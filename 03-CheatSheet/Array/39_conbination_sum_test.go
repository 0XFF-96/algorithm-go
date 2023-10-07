package array

import (
	"fmt"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T){
	candidates := []int{2, 3, 6, 7}
	// target := 7
	res := combinationSumV2(candidates, 7)
	t.Log(res)
}

// 因为全局变量，深拷贝的问题，出了问题
// 但是，为什么是 xxx ？
// 为什么最终结果是 [[7, 3, 3] | 7]

// 这里需要深拷贝的原因是， ans = append(ans, cur)
// 1- cur 是一个全局变量。
// 2- cur 之后还会发生多次变化。
// more-?

// 4 种不同解法。
// https://leetcode.com/problems/combination-sum/discuss/432145/Go-golang-0ms-four-solutions
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum.
//Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Combination Sum.
func combinationSumV2(candidates []int, target int) [][]int {
	var ans [][]int
	var cur []int
	sort.Ints(candidates)
	// dfs(candidates, target, 0, &cur, &ans)
	var dfs func(candidates []int, target int, s int)

	dfs = func(candidates []int, target int, s int){
		if target == 0 {
			fmt.Println("-----target----")
			fmt.Println(cur)
			fmt.Println("-------")
			ans = append(ans, append([]int{}, cur...))
			// ans = append(ans, append([]int{{, cur))
			fmt.Println("res", ans)
			fmt.Println("---------------")
			return
		}
		fmt.Println(cur)
		for i:=s; i < len(candidates); i++{
			if candidates[i] > target {
				break
			}
			cur = append(cur, candidates[i])
			fmt.Println(cur)
			// i 尤其重要
			// 为什么不是 i+1?
			dfs(candidates, target - candidates[i], i)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(candidates, target, 0)
	fmt.Println("----")
	fmt.Println(ans)
	fmt.Println("---")
	return ans
}

