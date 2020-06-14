package _18_pacal_s_triangle


func generate(numRows int) [][]int {
	var ret [][]int
	for i:=0; i<numRows;i++{
		tmp := make([]int, i+1)
		tmp[0] = 1
		ret = append(ret, tmp)

		for j:=1;j<=i;j++{
			if j == i {
				ret[i][j] = 1
			} else {
				ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
			}
		}
	}
	return ret
}