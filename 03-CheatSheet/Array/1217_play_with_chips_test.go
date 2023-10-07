package array


func minCostToMoveChips(chips []int) int {
	if len(chips) == 0 || len(chips) == 1{
		return 0
	}

	minCost := 1000000000

	for idx:=0;idx < len(chips)-1;idx++ {
		cost := 0
		j := idx+1
		for {
			if j == idx {
				break
			}
			steps := abs(j, idx)
			tmpCost := steps % 2
			cost += tmpCost
			j++
			j = j % len(chips)
		}
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

// 这么简单的思想
// 这道题目，是我想多了..
// 直接利用数学性质，就可以解决的问题。
func minCostToMoveChipsV2(chips []int) int {
	odd := 0
	even := 0
	for _, v := range chips {
		if v % 2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if odd < even {
		return odd
	}
	return even
}


//
// [1,2,2,2,2]
// 这种算是 1 ？
//func abs(a, b int) int {
//	if a > b {
//		return a - b
//	} else {
//		return b - a
//	}
//}
