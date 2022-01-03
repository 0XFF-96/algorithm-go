package _8_Subset



// 错误做法
func subsets(nums []int) [][]int {
	var ret [][]int
	var dfs func (n int, s int, ret []int)

	dfs = func(n int, s int, cur []int)  {
		// 实现 Cxx 的关键
		if len(cur) == n {
			// 有可能是深拷贝的问题
			ret = append(ret, cur)
			return
		}

		for i := s; i <= n; i ++{
			cur = append(cur, nums[i])
			dfs(n, i +1, cur)
			// pop
			cur = cur[:len(cur)-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		dfs(i, 0, []int{})
	}

	return ret
}
