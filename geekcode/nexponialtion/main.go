package nexponialtion

import (
	"math"

	"source-code-reading/go/src/testing"
)

// 快速幂运算
// 指数为低的情况
//


// 数值的 N 次方问题
// https://cloud.tencent.com/developer/article/1190761
// 包含相关的测试用例
// 输入的指数（exponent）小于1即是零和负数的时候怎么办

// 方法一：
// 1.全面考察指数的正负、底数是否为零等情况。
// * 2.写出指数的二进制表达，例如13表达为二进制1101。
// * 3.举例:10^1101 = 10^0001*10^0100*10^1000。
// * 4.通过&1和>>1来逐位读取1101，为1时将该位代表的乘数累乘到最终结果。

// 方法二
// 迭代快速幂算法...
// https://blog.csdn.net/hkdgjqr/article/details/5381028
func TestPower(t *testing.T){
	testCase := []struct {
		base float32
		exponent int
		expected float32
	}{
		{
			2, 3, 8,
		},
		{
			-2, 3, -8,
		},
		{
			2, -3, 0.125,
		},
		{
			2, 0, 1,
		},
		{
			0, 0, 1,
		},
		{
			0, 4, 0,
		},
		{
			0, -4, 0,
		},
		{
			0, 4, 0,
		},
	}
	for _, tt := range testCase {
		got := Power(tt.base, tt.exponent)
		t.Log(got, tt.expected)
	}
}



func Power(base float32, exponent int) (result float32) {
	if base == 0 && exponent < 0 {
		return 0
	}
	// 必须初始化一个值
	result = 1
	// base float32 不能直接用 == 判断
	absExponent := int(math.Abs(float64(exponent)))
	for i:=0; i < absExponent;i++{
		result *= base
	}
	if exponent < 0 {
		result = 1/ result
	}
	return result
}


// 二进制 1 的个数
// n & (n-1) 用于 + counter 来解决