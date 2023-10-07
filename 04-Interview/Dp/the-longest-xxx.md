- 最长括号匹配

```
func longestValidParentheses(s string) int {
  // 这个是否算有效括号，算！ ((())))
  // ())) + ')'

  if len(s) == 0 {
    return 0 
  }

  // 全部初始化为 0 
  dp := make([]int, len(s))
  ans := 0 
  for i := 1; i < len(s); i++{
    if s[i] == ')' {
      if s[i-1] == '('{
        tmp := 0 
        if i-2 > 0 {
          tmp = dp[i-2]
        } 
        dp[i] = 2 + tmp 
      } else  {

        idx := i - dp[i-1] - 1 
        if idx >= 0 && s[idx] == '(' {
                    tmp := 0 
        if idx - 1 >= 0 {
          tmp = dp[idx-1]
        }
        dp[i] = dp[i-1] + 2 + tmp 
        } else {
            dp[i] = 0 
        }

        // the test case is below:
        // ```
        // "()))"
        // '(())))'
        // '((())))'
        // ```
      }
    }
    ans = max(ans, dp[i])
  }
  return ans 
}
```


