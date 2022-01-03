package main 

// If you were only permitted to complete at most one transaction 
// (i.e., buy one and sell one share of the stock),

// import "math"
// func maxProfit(prices []int) int {
//     minPrice := math.MaxInt32
//     maxProfit := 0 
    
//     for i:=0;i < len(prices);i++{
//         if prices[i] < minPrice {
//             minPrice = prices[i]
//         } else if prices[i] - minPrice > maxProfit {
//             maxProfit = prices[i] - minPrice
//         }
//     }
//     return maxProfit
// }


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

https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/solution/

// 相关情况
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/solution/