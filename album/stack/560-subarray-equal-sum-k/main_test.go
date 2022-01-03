package main 

// 题目: https://leetcode.com/problems/subarray-sum-equals-k/solution/

// 没有 AC 
// 因为，没有举出反例子的定义。 
// 
func subarraySumV1(nums []int, k int) int {
    // 定义清楚： 【 continuous subarrays 】
    // 【 sequence array 】 
    // 【 subarrays 】
    // find total number 
    // if do't have ?
    // restriction [-1000, 1000]
    
    // continueous subarrays 可以不连续，但是需要组合。
    // 有多少个不同的子数组 == k 
    // 遍历子数组，
    // 判断是否相等。
    
    // [0,0,0,0,0,0,0,0,0,0]   = 55 ?
    // 例子反例子举一个。
    // dfs 来解决这个问题。 
    // 这里变成了一个搜索的问题。 
    totalNumber := 0
    for i := 0; i < len(nums); i++ {
        subArray := 0
        for j := i; j < len(nums); j++{
            subArray += nums[j]
            if subArray == k {
				totalNumber += 1
				
				// 这里 break 和 不 break 的区别很大。 
				// 是否正确答案？ 取决于此
				// [0,0,0,0,0,0,0,0,0,0]
                // break
            }
        }
    }
    return totalNumber
}


// 暴力方法, 不行
// 
func subarraySumV2(nums []int, k int) int {
    count := 0 
    
    // 计算 continue subarray 的基本方法
    for start:= 0; start < len(nums); start ++ {
        for end := start+1; end <= len(nums); end++{
            sum := 0
            for i := start; i < end; i ++ {
                sum += nums[i]
            }
            if sum == k {
                count++
            }
        }
    }
    return count
}

//  a cumulative sum array
// 这个技巧，用户快速算出 array[i: i+n] 子数组中的值，而省略一次 for 循环。
// 基本操作。
// in order to calculate the sum of elements lying between two indices, 
// we can subtract the cumulative sum corresponding to the two indices to obtain the sum directly, 
// instead of iterating over the subarray to obtain the sum.


// AC 了
// 但是还是不够快

// Runtime: 440 ms, faster than 6.38% of Go online submissions for Subarray Sum Equals K.
// Memory Usage: 5.9 MB, less than 73.33% of Go online submissions for Subarray Sum Equals K.
func subarraySumV3(nums []int, k int) int {
    count := 0 
    
    sum := make([]int, len(nums)+1)
    
    // sentinel number
    sum[0] = 0 
    for i:=1; i <= len(nums); i ++{
        sum[i] = sum[i-1] + nums[i-1]
    }
    
    // 计算 continue subarray 的基本方法
    for start:= 0; start < len(nums); start ++ {
        for end := start+1; end <= len(nums); end++{
            if sum[end] - sum[start] == k {
                count++
            }
        }
    }
    return count
}


// The idea behind this approach is as follows: If the cumulative sum(repreesnted by sum[i]sum[i] for sum upto i^{th}i th
// 	  index) upto two indices is the same, the sum of the elements lying in between those indices is zero. Extending the same thought further, if the cumulative sum upto two indices, say ii and jj is at a difference of kk i.e. if sum[i] - sum[j] = ksum[i]−sum[j]=k, the sum of elements lying between indices ii and jj is kk.


func subarraySum(nums []int, k int) int {
	// sum - k has occured already ?
   
   // Time complexity : O(n)O(n). The entire numsnums array is traversed only once.

   // Space complexity : O(n)O(n). Hashmap mapmap can contain upto nn distinct entries in the worst case.
   
   // 这种解法和 two-sum 有一定的类似。 
   // 值得参考学习。 
   
   
   // sum -k 意味着什么？
   // 为什么存的是 sum ?
   
   // 从如果，两个 prefix sum 相同， 中间的子数组的和一定为 0 
   // 延伸这个思想， 可以得出。
   // If the cumulative sum(repreesnted by sum[i]sum[i] for sum upto i^{th}i 
   // thindex) upto two indices is the same
   
   //  if sum[i] - sum[j] = ksum[i]−sum[j]=k, the sum of elements lying between indices ii and jj is kk.
   
   // sum[i] - k = sum[j]
   // 只要每次记录 sum 的情况即可以。
   // 而且需要cout
   
   count := 0 
   sum := 0 
   m := map[int]int{0:1}
   for i :=0; i < len(nums); i ++ {
	   sum += nums[i]
	   _, ok := m[sum-k]
	   if ok {
		   count += m[sum-k]
		  //  m[sum-k] += 1
	   } 
	   
	   if _, ok := m[sum]; ok {
		   m[sum] += 1
	   } else {
		   m[sum] = 1
	   }
   }
   
   // [1,2,3]
   // expected 3
   // got      1
   // fmt.Println(m)
   return count
}