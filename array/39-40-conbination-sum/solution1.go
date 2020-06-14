package  main

// 如何去重？
//
func conbinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil 
	}

	var res [][]int 
	var cur [][]int 

	var dfs func(candidates []int, target, curPos int) 

	dfs = func(candidates []int, target, curPos int) {
		if len(candidates) == 0 {
			return 
		}
		if target == 0 {
			// 坑: 这里必须要深拷贝
			// 因为引用的是函数全局变量
			res = append(res, append([]int{}, cur...))
		}
		for i:=curPos; i < len(candidates);i++{
			cur = append(cur, candidates[i])
			// 这里为什么是 i 
			// 是因为能够利用重复元素
			dfs(candidates, tareget-canidates[i], i)
			// backtracking
			// 弹出栈操作
			cur = cur[:len(cur)-1]
		}	
	}

	dfs(candiates, target, 0)
	return ans 
}