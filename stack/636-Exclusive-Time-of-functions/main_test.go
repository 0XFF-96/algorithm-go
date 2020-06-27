package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestExclusiveTime(t *testing.T){
	res := exclusiveTime(2, []string{
		"0:start:0","1:start:2","1:end:5","0:end:6",
	})

	t.Log(res)
}

// 不 AC 的题目
// 改进优化点， 单独写一个 parse log 函数
func exclusiveTime(n int, logs []string) []int {
    // logs in timestamp order 
    // [ [] ] 像一个括号，互相包围
    // [] [] []
    // [ [] [] [ [] ] ]
    // [ [ [] ] ]  
    // 计算这几种情况的区间计算 
    // 和括号匹配差不多。
    // 
    // each function has unique id 
    // 0 and n-1
    if n == 0 {
        return nil
    }
    res := make([]int, n)
    stack := []int{}
    s := strings.Split(logs[0], ":")
    
    id,_ := strconv.Atoi(s[0])
    stack = append(stack, id)

    // 从二号 log 开始
    i := 1
    time,_ := strconv.Atoi(s[2])
    for i < len(logs) {
        s = strings.Split(logs[i], ":")
        
        // 增加时间
        t,_ := strconv.Atoi(s[2])
        for time < t {
            res[stack[len(stack)-1]]++
            time++
        }
        if s[1] == "start" {
            id, _ = strconv.Atoi(s[0])
            stack = append(stack, id)
        } else {
            res[stack[len(stack)-1]]++
            time++
            stack = stack[:len(stack)-1]
        }
        i++
    }
    return res
}



// AC 版本。
// 比上一个版本增加了不少东西。
// 
func exclusiveTimeV2(n int, logs []string) []int {
    if n == 0 {
        return nil
    }
    res := make([]int, n)
    stack := []int{}
    s := strings.Split(logs[0], ":")
    id,_ := strconv.Atoi(s[0])
    stack = append(stack, id)
    i := 1
    prev,_ := strconv.Atoi(s[2])
    for i < len(logs) {
        // s = strings.Split(logs[i], ":")
        id, s , time := parseLog(logs[i])
        
        if s == "start" {
            // 这个判断非常重要
            if len(stack) != 0 {
                res[stack[len(stack)-1]] += time - prev
            }
            stack = append(stack, id)
            
            // 这里为什么不加一？
            // start 
            prev = time
        } else {
            res[stack[len(stack)-1]] += time - prev + 1 
            stack = stack[:len(stack)-1]
            
            // 这里为什么加 1 ？
            // end 状态
            prev = time + 1
        }
        i++
    }
    return res
}


func parseLog(log string)(int, string, int){
    strs := strings.Split(log, ":")
    id, _ := strconv.Atoi(strs[0])    
    time, _ := strconv.Atoi(strs[2])
    return id, strs[1], time
}