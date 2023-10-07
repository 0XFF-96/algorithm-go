package array



import "sort"
func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	origin := append([]int{}, arr...)
	sort.Ints(arr)


	// 有没有更好的办法，初始化这个变量呢？
	// 时间还不是最优的，
	// 消耗在哪里最大？
	//
	// 可以优化一个最后一个 for 循环，
	// 把它放在上一个 for 里面
	// 但好像不太好的样子。
	rank := 1
	m := map[int]int{}
	m[arr[0]] = 1
	for i:=1;i<len(arr);i++ {
		if arr[i] == arr[i-1] {
			continue
		}
		rank += 1
		m[arr[i]] = rank
	}

	var ret []int
	for _, v := range origin {
		ret = append(ret, m[v])
	}
	return ret
}
