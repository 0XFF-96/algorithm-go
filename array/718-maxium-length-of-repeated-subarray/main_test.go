package main 

// 基础版本
func findLengthV1(A []int, B []int) int {
    // 有给暴力的解法
    // A -> 所有子数组， 存一个 map, 
    // B -> 所有子数组,  搜索 map, 对比一下长度， 然后记录最大值 。 
    // -------------
    
    // 非暴力解法。
    // subarray 
    
    ans := 0 
    for i:=0; i < len(A);i++{
        for j:=0;j < len(B); j++{
			k := 0 
			
			// 优化点一
			// 大部分时间，这个条件都不会成立。
			// 所有可以用 hashmap 来优化一下。
            for i+k < len(A) && j+k < len(B) &&A[i+k] == B[j+k]  {
                k += 1
            }
            ans = max(ans, k)
        }
    }
    return ans
}

func max(i, j int) int {
    if i > j {
        return i
    } else {
        return j
    }
}



// 优化点一：
// only loop 
// https://leetcode.com/problems/maximum-length-of-repeated-subarray/submissions/
// 有一个 Bug 没有看出来。 
func findLength(A []int, B []int) int {   
    ans := 0 
    Bstarts := map[int][]int{}
    for j, y := range B {
        if _, ok := Bstarts[y];ok {
            Bstarts[y] = append(Bstarts[y], j)
        } else {
            Bstarts[y] = []int{j}  
        }
    }
    fmt.Println(Bstarts)
    
    for i, x := range A {
        if _, ok := Bstarts[x]; ok {
            for j := range Bstarts[x] {
                k := 0 
                for i+k < len(A) && j+k < len(B) && A[i+k] == B[j+k] {
                    k += 1
                }
                ans = max(ans, k)
            }
        }
    }
    return ans 
}