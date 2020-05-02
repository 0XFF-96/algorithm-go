package main

import (
	"testing"
)

func TestUniq(t *testing.T){
	number := [][]int{
		{
			0, 0,
		},
		{
			1, 1,
		},
		{
			0, 0,
		},
	}
	num := uniquePathsWithObstacles(number)
	t.Log(num)
}