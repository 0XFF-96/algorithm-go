
# Find candy swap 
- 888. Fair Candy Swap

```
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {

    //Restriction&Expectation,
    // Since they are friends, they would like to exchange one candy box each so that after the exchange, they both have the same total amount of candy
    
    // a  b 
    // a + i - j = b - i + j 
    // 找到 i, j 的过程
    alice, bob, aliceCandies := 0, 0, map[int]bool{}
    for _, candy := range aliceSizes { alice += candy; aliceCandies[candy] = true }
    for _, candy := range bobSizes { bob += candy }
    delta := (alice-bob)/2
    for _, b := range bobSizes { if aliceCandies[b+delta] { return []int{b+delta, b} } }
    return nil
}
```
