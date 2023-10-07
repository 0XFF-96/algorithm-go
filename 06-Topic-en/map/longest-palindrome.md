
# longest palindrome 

```
func longestPalindrome(s string) int {
    if len(s) == 0 {
        return 0 
    }

    count := map[byte]int{}
    length := len(s)

    for i :=0; i < length; i++{
        count[s[i]]++
    }

    ans := 0 
    for _, v := range count {
        ans += ( v / 2 ) * 2 
        if v % 2 == 1 && ans % 2 == 0 {
            ans++ 
        }
    }

    return ans 
}

```

