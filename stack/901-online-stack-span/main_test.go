type StockSpanner struct {
    priceStack []int
    spanStack []int
}


func Constructor() StockSpanner {
    return StockSpanner{make([]int, 0), make([]int, 0)}
}


func (this *StockSpanner) Next(price int) int {
    result := 1
    for len(this.priceStack) > 0 && this.priceStack[len(this.priceStack) - 1] <= price {
        result += this.spanStack[len(this.spanStack) - 1]
        this.priceStack = this.priceStack[:len(this.priceStack) - 1]
        this.spanStack = this.spanStack[:len(this.spanStack) - 1]
    }
    this.priceStack = append(this.priceStack, price)
    this.spanStack = append(this.spanStack, result)
    return result
}


