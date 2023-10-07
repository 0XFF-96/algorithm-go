# Find the target indices 

```

func targetIndices(nums []int, target int) []int {
    result := make([]int, 0 )
    startIndex := 0 
    total := 0 
    
    for _, num := range nums {
        if num < target {
            startIndex++
        } else if num == target {
            total++
        }
    }

    for i := 0; i < total; i ++ {
        result = append(result, startIndex + i)
    }
    return result 
}

```
