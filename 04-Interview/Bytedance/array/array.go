

// 缺失的第一个正数
// 

// 暴力解法的思路
// 我们可以将数组所有的数放入哈希表，随后从 11 开始依次枚举正整数，并判断其是否在哈希表中；
// 我们可以从 11 开始依次枚举正整数，并遍历数组，判断其是否在数组中。 

// hash 计数：需要额外的空间。所有出现的数字，hash 值 + 1。
// 第二次从 1 开始查找 hash 表，找不到就是它了。
// 但是这种方案不符合题意，而且 hash 计数这个方案效率也不是最优的

// 缺点是：需要用哈希表， hashmap , 有一定的存储空间。 哈希表的本质就是一个数组/链表 + 哈希函数。 
// 怎么在这个基础上， 去除对 hashmap 的一个依赖。 

// 题目的条件限制。 


// 具体的思路是：
// 遍历一次数组把大于等于1的和小于数组大小的值放到原数组对应位置，
// 然后再遍历一次数组查当前下标是否和值对应，如果不对应那这个下标就是答案，
// 否则遍历完都没出现那么答案就是数组长度加1

// 思路：
// 

// 我们对数组进行遍历，对于遍历到的数 xx，如果它在 [1, N][1,N] 的范围内，
// 那么就将数组中的第 x-1x−1 个位置（注意：数组下标从 00 开始）打上「标记」。
// 在遍历结束之后，如果所有的位置都被打上了标记，那么答案是 N+1N+1，
// 否则答案是最小的没有打上标记的位置加 1 
// https://leetcode-cn.com/problems/first-missing-positive/solution/que-shi-de-di-yi-ge-zheng-shu-by-leetcode-solution/。 


func firstMissingPositive(nums []int) int {
    n := len(nums)
    for i := 0; i < n; i++ {
        if nums[i] <= 0 {
            nums[i] = n + 1
        }
    }
    for i := 0; i < n; i++ {
        num := abs(nums[i])
        if num <= n {
            fmt.Println(num-1)
            nums[num - 1] = -abs(nums[num - 1])
        }
    }
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            return i + 1
        }
    }
    return n + 1
}
