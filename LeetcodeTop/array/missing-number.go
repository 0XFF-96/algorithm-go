
// 丢失数字
// 


func missingNumber(nums []int) (xor int) {
    for i, num := range nums {
        xor ^= i ^ num
    }
    return xor ^ len(nums)
}



// 数学公式， 
// 将从 00 到 nn 的全部整数之和记为 \textit{total}total，根据高斯求和公式，有：
// 将数组 \textit{nums}nums 的元素之和记为 \textit{arrSum}arrSum，则 \textit{arrSum}arrSum 比 \textit{total}total 少了丢失的一个数字，因此丢失的数字即为 \textit{total}total 与 \textit{arrSum}arrSum 之差。
// 给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
// [总数]

// 符合高斯公式。 
func missingNumber(nums []int) int {
    n := len(nums)
    total := n * (n + 1) / 2
    arrSum := 0
    for _, num := range nums {
        arrSum += num
    }
    return total - arrSum
}