package main 

// 怎么进行 AC 处理。
// Brute Force [Time Limit Exceeded]
// AC 了，但是不通过。
func arrayNesting(nums []int) int {
    // nest array 
    // A[A[i]].... 
    // 用处在哪里？
    
    // dfs 
    // 特点 0-n 
    // 数组长度固定
    max := 0 
    for i := 0; i < len(nums); i ++ {
        set := map[int]struct{}{}
        count := 0
        prev := nums[i]
        // fmt.Println(prev)
		// while 的逻辑
		
		// 如何简化这个循环逻辑？
		// 这个逻辑在做什么？
		// 有没有可能 early retrun ?
		// 
        for {
            // 有没有可能死循环？
            // 
            _, ok := set[prev]
            if ok {
                break
            }
            set[prev] = struct{}{}
            count++
            if count > max {
                max = count
            }
			prev = nums[prev]
			
			// 这个 prev 变量起得并不好。 
			// 很容易混淆，改的时候，会容易出错。
            // fmt.Println(prev)
            // set[prev] = struct{}{}
        }
    }
    return max 
}



// 优化方法。
// Using Visited Array [Accepted]
// https://leetcode.com/problems/array-nesting/solution/


// 在哪里程序进行了重复计算？
// wrost case 是哪个？
// 所有的元素都被 add set, n 平方的计算量。 
// 你能否构造一个这样的解出来？
// 知道为什么上面的算法不行，才能写出一个符合要求的正确解法。
// 


// 代码的放置顺序。！
func arrayNesting(nums []int) int {
    visited := make([]bool, len(nums))
    res := 0 
    
    for i:=0; i < len(nums); i++{
        if (!visited[i]) {
            start := nums[i]
            count := 0
            for {
                start = nums[start]
                count++
                visited[start] = true
                
                if count > res {
                    res = count
                }
                
                if start == nums[i] {
                    break
                }
            }   
        }
    }
    return res
}