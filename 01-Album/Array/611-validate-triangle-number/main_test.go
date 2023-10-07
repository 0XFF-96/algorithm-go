package main 

func triangleNumber(nums []int) int {
    // sorting  ?
    // early stop 
    // combination , 3
    // dfs 
    // 怎么写回溯的代码。 append , dfs, remove 
    // 三角形状不等公式。 两边之和要大于第三边。
    
    // Using Binary Search
    // a ≤ b ≤ c, we need not check all the three inequalities
    
    // shows how Binary Search can be used to find the right limit for a simple example.
    sort.Ints(nums)
    count := 0 
    for i := 0; i < len(nums) -2; i++ {
        k := i + 2
        for j := i + 1; j < len(nums) -1 && nums[i] != 0; j ++ {
            k = binarySearch(nums, k, len(nums)-1, nums[i] + nums[j])
            count += k - j - 1;
        }
    }
    return count 
}

func binarySearch(nums []int, l , r , x int) int {
    for r >=l && r < len(nums) {
        mid := (l + r) / 2 
        if nums[mid] >= x {
            r = mid -1
        } else {
            l = mid + 1
        }
    }
    // 
    return l
}
