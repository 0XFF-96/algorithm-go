package main

func mctFromLeafValues(arr []int) int {
    // 哈夫曼树相关算法的实现
    // 2 4 6 7    27 46 
    // 8 13       9   10
    // 21         19
    // 
    // 这是一种构建哈夫曼树🌲的相关算法。
    // 排序数组
    // 取最小的两个元素，构建节点
    // 把结果放回到数组里面
    // 重新排序
    // 然后，重复这个过程。直到数组中的元素只有 1 个。 
    res := 0 
    stack := []int{10000}
    for i := 0; i < len(arr); i ++ {
        for stack[len(stack)-1] <= arr[i] {
            mid := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res += mid * min(stack[len(stack)-1], arr[i])
        }
        stack = append(stack, arr[i])
    }
    
    for len(stack) > 2 {
        pop := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res += stack[len(stack)-1] * pop
        // 这里之前用
        // stack = stack[:len(stack)-2]
        // 导致有一个测试用例，不过
    }
    return res 
}

func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}

//     def mctFromLeafValues(self, A):
//         res = 0
//         while len(A) > 1:
//             i = A.index(min(A))
//             res += min(A[i - 1:i] + A[i + 1:i + 2]) * A.pop(i)
//         return res


// 使用栈
    // public int mctFromLeafValues(int[] A) {
    //     int res = 0;
    //     Stack<Integer> stack = new Stack<>();
    //     stack.push(Integer.MAX_VALUE);
    //     for (int a : A) {
    //         while (stack.peek() <= a) {
    //             int mid = stack.pop();
    //             res += mid * Math.min(stack.peek(), a);
    //         }
    //         stack.push(a);
    //     }
    //     while (stack.size() > 2) {
    //         res += stack.pop() * stack.peek();
    //     }
    //     return res;
    // }