
# 搜索二维矩阵

- 方法一： 根据矩阵的相关特性， 开始进行相关搜索。 

```

func searchMatrix(matrix [][]int, target int) bool {
    var r, c = len(matrix), len(matrix[0])
	for i,j := r - 1, 0;i >=0 && j < c; {
		if matrix[i][j] == target{
			return true
		}else if matrix[i][j] < target{
			j++
		}else{
			i--
		}
	}
	return false
}

```


```

func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])

    left, right := 0, m * n - 1 
    for left <= right {
        // 重点是把二维矩阵，扯平的这个操作比较有意思。
        var mid = left + (right - left) / 2 
        curRow := mid / n 
        curCol := mid % n 
        var num = matrix[ curRow ][ curCol ] 
        if num == target {
            return true 
        }

        if num < target {
            left = mid + 1 
        } else {
            right = mid - 1 
        }
    }
    return false 
}

```


