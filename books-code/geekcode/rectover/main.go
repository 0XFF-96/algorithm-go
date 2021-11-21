package rectover

import (
	"testing"
)

// 跳台阶问题
// 可以等价于斐波那契数列的问题
// dp(n) = dp(n-1) + dp(n-2) 的等价问题
// 用迭代来解决

// 矩形覆盖
func TestRectCover(t *testing.T){
	testCase := []struct{
		input int
		ouput int
	}{
		{
			1,
			1,
		},
		{
			4,
			5,
		},
		{
			5,
			8,
		},
	}

	for _, tt := range testCase{
		res := rectCover(tt.input)
		_ = res
		// require.Equal(t, tt.ouput, res)
	}

}

func rectCover(target int) int {
	if target <= 2 {
		return target
	}
	pre2 := 1
	pre1 := 2
	cur := 0
	if target == 0 {
		return 0
	}
	for i:=3; i <= target; i++{
		cur = pre1 + pre2
		pre2 = pre1
		pre1 = cur
	}
	return pre1
}
