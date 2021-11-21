package array



// 一次 AC
// 双 100 +
//
func largeGroupPositions(S string) [][]int {
	if len(S) == 0 {
		return nil
	}
	// one-pass
	// 1、>= 3
	// 2、lexicographic order
	// 3、how to sort a string by alphetica order ?
	//  lexicographic 字典序 ?

	startIdx := 0
	endIdx := 0

	var ret [][]int
	for i := 1; i < len(S); i++ {
		if S[i-1] == S[i] {
			endIdx = i
		} else {
			if endIdx - startIdx >= 2 {
				ret = append(ret, []int{startIdx, endIdx})
			}
			startIdx = i
			endIdx = i
		}

	}

	// for case like aaaaaaaaa
	// 这里的重复代码，有没有什么办法可以消除？
	// 最后一次循环的时候，判断一下.
	if endIdx - startIdx >= 2 {
		ret = append(ret, []int{startIdx, endIdx})
	}

	return ret
}

// ......
// ......
// https://leetcode.com/articles/positions-of-large-groups/