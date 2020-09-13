package main 

func findShortestSubArray(nums []int) int {
    tmp := make(map[int][]int)
	degree := 0
	// 最大
	subArr := 50000
    
    for idx, n := range nums {
        // 这里为什么不需要初始化？
        tmp[n] = append(tmp[n], idx) 
        if len(tmp[n]) > degree {
            degree = len(tmp[n])
        }
    }
    
    for k, v := range tmp {
        if len(v) == degree {
            
            // 这里为什么需要加 1 ？
            l := v[len(tmp[k])-1] - v[0] + 1
            if l < subArr {
                subArr = l
            }
        }
    }
    return subArr