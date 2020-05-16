package jump_floor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// 变态跳台阶问题
// https://www.nowcoder.com/profile/3754286/codeBookDetail?submissionId=13753640
// 算法实现如下。
func TestJumpFloor(t *testing.T){
	tests := []struct{
		intput int
		output int
	}{
		{
			0,
			0,
		},
		{
			1,
			1,
		},
		{
			2,
			2,
		},
		{
			3,
			4,
		},
		{
			4,
			8,
		},
	}

	for _, tt := range tests{
		require.Equal(t, tt.output, jumpNFloor(tt.intput))
	}
	// PS:
	// 使用位移更快
	// 2^(n-1)可以用位移操作进行，更快
}

func jumpNFloor(stair int) int {
	if stair == 0 {
		return 0
	}
	if stair == 1 {
		return 1
	}
	if stair == 2 {
		return 2
	}
	return 2 * jumpNFloor(stair-1)
}
