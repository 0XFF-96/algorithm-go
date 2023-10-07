
# min array 

```

func minArray(numbers []int) int {
    i, j := 0, len(numbers) -1 

    for i < j {
        m := ( i + j ) / 2 
        if numbers[m] > numbers[j] {
            i = m + 1 // 左有序
        } else if numbers[m] < numbers[j] {
            j = m 
        } else {
            j -= 1 // 收缩边界
        }
    }
    return numbers[i]
}

```


