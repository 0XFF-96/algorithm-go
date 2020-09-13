package main 
func maxProfit(prices []int, fee int) int {    
    // Consider the first K stock prices. 
    // At the end, the only legal states are that you don't own a share of stock, 
    // or that you do. Calculate the most profit 
    // you could have under each of these two cases.
    cash, hold := 0, -prices[0]
    for i:=1; i < len(prices); i ++ {
        cash = max(cash, hold + prices[i]-fee)
        hold = max(hold, cash - prices[i])
    }
    return cash 
    
    // cash 
    // the maximum profit we could have if we did not have a share of stock
    
    // hold 
    // the maximum profit we could have if we owned a share of stock.
    
    // sell 
    //. cash = max(cash, hold + prices[i] - fee)
    
    // buy 
    // hold = max(hold, cash - prices[i]
}

func max(i , k int) int {
    if i > k {
        return i
    }   else {
        return k
    }
}