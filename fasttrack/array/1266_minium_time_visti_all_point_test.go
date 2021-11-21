package array


// 一开始想多了
// 其实就是，从一个点走到另一个点，有哪几种方法的问题。并且累计计算就好。
// 两点之间的距离和什么有关系。
func minTimeToVisitAllPoints(points [][]int) int {
	// You have to visit the points in the same order as they appear in the array.
	// 1、To walk from point A to point B there will be an optimal strategy to walk ?
	// 2、Advance in diagonal as possible then after that go in straight line.
	// 3、Repeat the process until visiting all the points.

	var dx, dy, result int
	var p0, p1 []int

	for i := 0; i < len(points)-1; i++ {
		p0, p1 = points[i], points[i+1]
		dx, dy = p0[0] - p1[0], p0[1] - p1[1]

		if dx < 0 {
			dx *= -1
		}

		if dy < 0 {
			dy *= -1
		}

		if dx >= dy {
			result += dx
		} else {
			result += dy
		}
	}
	return result
}
