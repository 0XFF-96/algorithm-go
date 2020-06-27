// https://www.youtube.com/watch?v=tiAaVfMcL9w



func scoreOfParentheses(S string) int {
    // score 
    if len(S) == 0 {
        return 0
    }
    return score(S, 0, len(S)-1)
}

func score(s string, l, r int) int {
    if r - l == 1 {
        return 1 // "()"
    }
    
    b := 0 
    for i:=l; i < r; i++{
        if string(s[i]) == "(" {
           b++ 
        } else {
            b--
        }
        
        if b == 0 {
            return score(s, l, i) + score(s, i+1, r)
        }
    }
    
    // 脱一层括号
    return 2 * score(s, l+1, r-1)
}

// 这里有一个非常厉害的转化思想
// 学一下
func scoreOfParenthesesV2(S string) int {
    // score 
    if len(S) == 0 {
        return 0
    }

    // the key idea is 
    // to convert the original string into 
    // a set of full balanced stringsf 
    
    ans := 0 
    d := -1 
    
    for i:=0; i < len(S); i++ {
        if string(S[i]) == "(" {
            d += 1
        } else {
            d -= 1
        }
        
        // 这个判断是什么意思？
        // key idea 的转化方法
        if string(S[i]) == "(" && string(S[i+1]) == ")" {
            // 移位操作
            ans += 1 << d
        }
    }
    return ans
}