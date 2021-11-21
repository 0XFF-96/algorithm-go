package array

import (
	"testing"
)

func TestImageSmoother(t *testing.T){


	// 通不过 test case [[2,3,4],[5,6,7],[8,9,10],[11,12,13],[14,15,16]]
	// 有一行代码错了...
	//
	ans := imageSmoother([][]int{{1,1, 1},{1,0,1},{1,1,1}})
	t.Log(ans)

	t.Log(3/4)
}

func imageSmoother(M [][]int) [][]int {
	row := len(M)
	col := len(M[0])

	// ans = make([][]int)
	// 必须是 const
	// golang 怎么初始化呢？

	// 初始化
	ans := make([][]int,row)
	for i:=0;i<row;i++{
		// ans[i] = append(ans[i], make([]int, col)...)
		ans[i] = make([]int, col)
	}
	// 初始化
	// fmt.Println(len(ans[0]))
	// 遍历二维数组
	// 的一个基本方法
	for r:=0;r < row; r++{
		for c:=0;c<col;c++{
			count := 0

			// i, j
			for nr := r-1; nr <= r+1; nr++ {
				for nc := c-1; nc <= c+1; nc++ {
					if 0 <= nr && nr < row && 0 <= nc && nc < col {
						// fmt.Println(M[nr][nc])
						ans[r][c] += M[nr][nc]
						// fmt.Println(ans[i][j])
						count++
					}
				}
			}
			// fmt.Println(ans[i][j])
			ans[r][c] /= count
		}
	}
	return ans
}
