package main 
// 基本解法
// 为什么这个不用 seen ?
func maxAreaOfIsland(grid [][]int) int {
    maxArea := 0
    for y, maxY := 0, len(grid); y < maxY; y++ {
        for x, maxX := 0, len(grid[y]); x < maxX; x++ {
            if grid[y][x] > 0 {
                if sum := sumIslandArea(grid, y, x); sum > maxArea {
                    maxArea = sum
                }
            }
        }
    }
    return maxArea
}

func sumIslandArea(grid [][]int, y int, x int) int {
	// 注意初始值
	sum := 1

	// 标记为 0 
	// 是为了省略 seen 而设计的
	// 厉害
	// prevent repeated visits 
    grid[y][x] = 0

    maxY, maxX := len(grid), len(grid[y])
    if y > 0 && grid[y-1][x] > 0 {
        sum += sumIslandArea(grid, y-1, x)
    }
    if x < maxX-1 && grid[y][x+1] > 0 {
        sum += sumIslandArea(grid, y, x+1)
    }
    if y < maxY-1 && grid[y+1][x] > 0 {
        sum += sumIslandArea(grid, y+1, x)
    }
    if x > 0 && grid[y][x-1] > 0 {
        sum += sumIslandArea(grid, y, x-1)
    }
    return sum
}