package main 


// k ,  absolute difference is k
//  
func findPairs(nums []int, K int) int {
    if K < 0 {
        return 0
    }
    m := make(map[int]int)
    for _, v := range nums {
        m[v]++
    }
    var out int
    for k, _ := range m {
        v2 := m[k+K]
        if (K == 0 && v2 == 1) || v2 == 0 {
            continue
        }
        out++
    }
    return out
}

