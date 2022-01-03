


func longestWPI(hours []int) int {
    // 读懂题目了 
   //  6， 6， 9， 9， 8， 6， 6， 6
   //  4 
   // 6, 6, 9, 9, 8, 6, 6, 9, 6 
   // 7
    // Make a new array A of +1/-1s corresponding to if hours[i] is > 8 or not. 
    // The goal is to find the longest subarray with positive sum
    // perfix sum !
    //  the smallest i < j with PrefixSum[i] + 1 == PrefixSum[j]
    
    
    // test case 
    // 
    // [9,9,9,6,0,6,6,9]
    // output 5 
    res := 0 
    score := 0 
    n := len(hours)
    seen := map[int]int{}   
    for i:=0; i < n; i++{
        if hours[i] > 8 {
            score +=1
        } else {
            score -= 1
        }
        if score > 0 {
            res = i + 1
        } else {
            // put if absent
            
            // 这两要好好研究一下
            if _, ok := seen[score]; !ok {
                seen[score] = i
            }
            
            if value, ok := seen[score-1]; ok {
                res = Max(res, i - value)
            }
        }
    }
    return res   
}


func Max(a, b int) int {
    if a > b {
        return a 
    } else {
        return b
    }
}



