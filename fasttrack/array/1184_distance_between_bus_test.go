package array


func distanceBetweenBusStops(distance []int, start int, destination int) int {
	cur1 := 0
	i := start
	// 怎么判断是否到了同一位置？
	for {
		if i == destination {
			break
		}
		cur1 += distance[i]
		i++
		i = i % (len(distance) )
	}

	cur2 := 0
	j := destination
	for {
		if j == start {
			break
		}
		cur2 += distance[j]
		j++

		// % 的是 len(distance)
		// 还是 distance + 1
		// 这很重要好不？
		j = j % (len(distance) )
	}
	if cur1 > cur2 {
		return cur2
	} else {
		return cur1
	}
}
