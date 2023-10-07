package main 


// 自己想的，但是有一部分功能没有通过。
type MyCalendar struct {
    c map[int]struct{}
}


func Constructor() MyCalendar {
    return MyCalendar{
        c: map[int]struct{}{},
    }
}

// 如何快速判断，一个点是否在一个数组内？
// 一个点，是否在一个线段内？
// 10， 20 
// 15， 25 
// 20 ， 30 

// 暴力方法， 用 map 存两条线段之间点中间点，
// 然后 For 循环，遍历加入点，但是这两种方法都是 O(n) 时间的。 
// O(1) 方法

//  De Morgan's laws
func (this *MyCalendar) Book(start int, end int) bool {
    if start == end || start > end {
        return false
    }   
    
    for i:= start+1; i < end; i++{
        if _, ok := this.c[i]; ok {
            return false
        }
    }
    
    for i:= start+1; i < end; i++{
        this.c[i] = struct{}{}
    }
    return true
}



// 二维数组


type MyCalendar struct {
    Activity [][]int
}


func Constructor() MyCalendar {
    return MyCalendar{}
}


func (this *MyCalendar) Book(start int, end int) bool {
    for _, v := range this.Activity {
        if v[0] < end && start < v[1] {
            return false
        }
    }
    this.Activity = append(this.Activity, []int{start, end})
    return true
}