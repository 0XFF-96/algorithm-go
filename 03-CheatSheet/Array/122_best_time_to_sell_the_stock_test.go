package array


// 这次的不点是
// 能够无限买卖。
// 所以从后遍历数组。

func maxProfitV3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	sumProfit :=0
	for i:=len(prices)-1;i>0;i--{
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			sumProfit += tmp
		}
	}
	return sumProfit
}


