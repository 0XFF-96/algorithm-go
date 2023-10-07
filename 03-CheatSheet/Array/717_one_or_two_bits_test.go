package array


func isOneBitCharacter(bits []int) bool {
	i := 0

	// 为什么需要减去 1
	for i < len(bits)-1 {
		i += bits[i] + 1
	}
	return i == len(bits) -1
}

// 不懂
// greedy solution 是什么意思？
// 为社么需要减去 2 ？
// https://leetcode.com/problems/1-bit-and-2-bit-characters/solution/
