package array


// https://leetcode.com/problems/fibonacci-number/
// 需要用尽 5 种方法解题
//


func fib(N int) int {
	if N < 1 {
		return 0
	}

	// 递归基
	f0 := 0
	f1 := 1
	fn := 1
	for i := 1; i <= N-1; i++{
		fn = f0 + f1
		f0 = f1
		f1 = fn
	}
	return fn
}