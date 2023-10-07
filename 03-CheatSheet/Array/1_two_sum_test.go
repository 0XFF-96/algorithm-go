package array


// 需要写出三种解法
// 优化的思想：固定某个变动参数。
// one-pass solution
func twoSum(nums []int, target int) []int {
	if len(nums) ==0 {
		return nil
	}
	vMap := map[int]int{}
	for i:=0;i < len(nums);i++{
		complement := target - nums[i]
		if idx, ok := vMap[complement];ok{
			return []int{idx,i}
		}
		vMap[nums[i]] = i
	}
	return nil
}

// solution 2 


// two-sum 的优化版本，如果数组有序的情况下，应该怎么办？
int[] twoSum(int[] nums, int target) {
    int left = 0, right = nums.length - 1;
    while (left < right) {
        int sum = nums[left] + nums[right];
        if (sum == target) {
            return new int[]{left, right};
        } else if (sum < target) {
            left++; // 让 sum 大一点
        } else if (sum > target) {
            right--; // 让 sum 小一点
        }
    }
    // 不存在这样两个数
    return new int[]{-1, -1};
}

