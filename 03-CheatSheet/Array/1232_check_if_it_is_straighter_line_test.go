package array

func checkStraightLine(coordinates [][]int) bool {
	// 斜率一样
	// gradient
	// we have the same slope
	// 两点确定一条直线
	slope := 1000000000.0
	for i:=0;i < len(coordinates) -1; i++{
		p1 := coordinates[i]
		p2 := coordinates[i+1]
		x := p1[0] - p2[0]
		y := p1[1] - p2[1]
		// abs ?
		tmpSlope := float64(y)/float64(x)
		if i == 0 {
			slope = tmpSlope
		}
		if tmpSlope != slope {
			return false
		}
	}
	return true
}
