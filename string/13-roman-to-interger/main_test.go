func romanToInt(s string) int {
    var result, curr, prev int
    for i := 0; i < len(s); i++ {
        curr = table(s[i])
        result += curr
        if prev < curr {
            // 重点语句
            result -= 2 * prev
        }
        prev = curr
    }
    return result
}

func table(c byte) int {
    switch c {
    case 'I': return 1
    case 'V': return 5
    case 'X': return 10
    case 'L': return 50
    case 'C': return 100
    case 'D': return 500
    case 'M': return 1000
    }
    return 0
}