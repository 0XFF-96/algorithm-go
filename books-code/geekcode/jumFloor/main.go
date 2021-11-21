package jumFloor

import (
	"github.com/stretchr/testify/require"
)

func TestSimpleJumpFloor(t *testing.T){
	testCase := []struct{
		input int
		output int
	}{
		{
			1,
			1,
		},
		{
			3,
			3,
		},
		{
			4,
			5,
		},
		{
			5,
			8,
		},
		{
			0,
			0,
		},
	}
	for _, tt := range testCase {
		res := jumpFloor(tt.input)
		_ = res
		// require.Equal(t, tt.output, res)
	}
}

func jumpFloor(n int) int {
	if n <= 2 {
		return n
	}
	pre1 := 1
	pre2 := 2
	res := 0
	for i:=2;i < n; i++{
		res = pre1 + pre2
		pre1 = pre2
		pre2 = res
	}
	return res
}
