# roman symbol 


```

// func romanToInt(s string) int {
//     // Problem is simpler to solve by working the string from back to front and using a map.
//     // 首先建立一个HashMap来映射符号和值，然后对字符串从左到右来，如果当前字符代表的值不小于其右边，就加上该值；否则就减去该值。以此类推到最左边的数
//     // 

// }

var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
    n := len(s)
    for i := range s {
        value := symbolValues[s[i]]
        if i < n-1 && value < symbolValues[s[i+1]] {
            ans -= value
        } else {
            ans += value
        }
    }
    return
}

```




