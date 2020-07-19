package main 

func findDuplicates(nums []int) []int {
    // 位运算
    // 找出出现两次的元素
    // 1 ≤ a[i] ≤ n (n = size of array)
    
}
    // when find a number i, flip the number at position i-1 to negative. 
    // if the number at position i-1 is already negative, i is the number that occurs twice.
    
    // public List<Integer> findDuplicates(int[] nums) {
    //     List<Integer> res = new ArrayList<>();
    //     for (int i = 0; i < nums.length; ++i) {
    //         int index = Math.abs(nums[i])-1;
    //         if (nums[index] < 0)
    //             res.add(Math.abs(index+1));
    //         nums[index] = -nums[index];
    //     }
    //     return res;
	// }
	
import (
		"math"
)

func findDuplicates(nums []int) []int {
	i := 0
	var res []int
		
	for i < (len(nums)) {
		index := int(math.Abs(float64(nums[i])))
		if nums[index-1] < 0 { 
			res = append(res, index)
		} else {  
				nums[index-1] = -nums[index-1]    
		}
		i++
	}
	return res
}