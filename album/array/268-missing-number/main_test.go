func TestMissing(t *testing.T) {
	missingNumber([]int{0, 1, 3, 5})
}


func missingNumber(nums []int) int {
    // sum up and subtract it with the largest
    // 等差数列公式，( n * (n + 1) ) / 2    
    
    var sum int
    for _, n := range nums {
        sum += n 
    }
    s := len(nums)
    r := (s * (s + 1)) /2 
    return r - sum
}


// 位运算的神奇之处
// =4∧(0∧0)∧(1∧1)∧(2∧3)∧(3∧4)
// =(4∧4)∧(0∧0)∧(1∧1)∧(3∧3)∧2
// =0∧0∧0∧0∧2
// =2
func missingNumberV2(nums []int) int {
    // sum up and subtract it with the largest
    // 等差数列公式，( n * (n + 1) ) / 2    
    
    
    // xor 的特点,
    // 异或
    // 按位与
    // 按位或
    missing := len(nums)
    
    for i, n := range nums {
        missing ^= i ^ n
    }
    return missing
}