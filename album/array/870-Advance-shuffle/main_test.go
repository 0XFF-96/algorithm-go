package main

type value_index struct {
    value int
    index int
}
func advantageCount(A []int, B []int) []int {
    // return any permutaion of A 
    // Why is this solution so complicated? Can't you sort A first, then for each element of B do an upper            bound of that element in A. If it exists use that, otherwise, use the smallest element in A instead.       // Either way, erase from A the element that you chose
    sort.Ints(A)
    n := len(A)
    b_val_idx := make([]value_index, n)
    for i, val := range B {
        b_val_idx[i].value = val
        b_val_idx[i].index = i
    }
    sort.Slice(b_val_idx[:], func(i, j int) bool{
        return b_val_idx[i].value > b_val_idx[j].value
    })
    lo, hi := 0, n-1
    res := make([]int, n)
    for _, v := range b_val_idx {
        idx, val := v.index, v.value
        if A[hi] > val {
            res[idx] = A[hi]
            hi--
        } else {
            res[idx] = A[lo]
            lo++
        }
    }
    return res
}