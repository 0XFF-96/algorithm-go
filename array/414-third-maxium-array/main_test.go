package main 

import "math"

func thirdMax(nums []int) int {
    save := []int{int(math.Inf(-1)),int(math.Inf(-1)),int(math.Inf(-1))}
    
    for _, v := range nums {
        if v != save[0] && v != save[1] && v != save[2] {
            if v > save[0] {
                save = []int{v, save[0], save[1]}
            }else if  v > save[1] {
                save = []int{save[0], v, save[1]}
            }else if v > save[2] {
                save[2] = v
            }            
        }
    }
    if save[2] == int(math.Inf(-1)){
		return save[0]
	}
    return save[2]
}