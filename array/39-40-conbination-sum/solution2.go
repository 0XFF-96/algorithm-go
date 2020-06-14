package main

func combinationSum(candidates []int, target int) [][]int {
    r := [][]int{}
    for i, c := range candidates {
        if c == target {
            r = append(r, []int{c})
        } else if c < target {
            for _, rr := range combinationSum(candidates[i:], target - c) {
                r = append(r, append(rr, c))
            }
        }
    }
    return r
}