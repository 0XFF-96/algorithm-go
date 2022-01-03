package main
import "sort"

func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return nil 
    }
    sort.Slice(intervals, func(i, j int) bool{
        return intervals[i][0] < intervals[j][0]
    })
    
    var res [][]int
    last := intervals[0]
    // intervals = append(intervals, last)
    for _, cur := range intervals {
        if cur[0] <= last[1] {
            if cur[1] > last[1] {
                last[1] = cur[1]
            }
        } else {
            res = append(res, last)
            last = cur
        }
    }
    
    // 直接合并最后一个
    res = append(res, last)
    return res        
}