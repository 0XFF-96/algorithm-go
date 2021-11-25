package array

// solution test:https://leetcode.com/problems/3sum-closest/discuss/7872/Java-solution-with-O(n2)-for-reference
// kind of
//

// // https://leetcode.com/problems/sum-of-nodes-with-even-valued-grandparent/discuss/535783/Go-golang-BFS-and-DFS-solutions
// 这种方法清晰一点
//
//func threeSumClosest(nums []int, target int) int {
//	sort.Ints(nums)
//	minDist := 1<<30
//	ans := 0
//	for i, v := range nums {
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		l, r := i+1, len(nums)-1
//		for l < r {
//			dist := v + nums[l] + nums[r]
//			if dist == target {
//				return target
//			}
//
//			if abs(target-dist) < minDist {
//				minDist = abs(target-dist)
//				ans = dist
//			}
//
//			if dist < target {
//				l++
//			} else {
//				r--
//			}
//		}
//	}
//	return ans
//}