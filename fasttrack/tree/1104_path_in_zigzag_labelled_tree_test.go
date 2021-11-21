package tree


// https://www.youtube.com/watch?v=YsLko6sSKh8
// https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree/
// zigzag travel tree


// 如何从一个孩子的节点得知 parent path ?
//


// 难以学到到骚操作
// 将 label 看作 二进制数字。
// https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree/discuss/323848/Golang-O(log-n)-with-detail-explanation
func pathInZigZagTree(label int) []int {
	ans := []int{label}

	for label > 1 {
		label = convert(label)
		ans = append(ans, label)
	}
	for i, j :=0,len(ans) -1;i < j;i,j = i+1, j-1{
		ans[i], ans[j] = ans[j], ans[i]
	}

	return ans
}


func convert(n int) int {
	tb := uint(bits(n))
	return symmetric(n, tb) >> 1
}

func symmetric(n int, tb uint) int {
	return 1 << tb + 1 <<(tb -1) - 1 -n
}

func bits(n int) int {
	r := 0
	for n > 0 {
		r++
		n >>= 1
	}
	return r
}