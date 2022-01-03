
### similar question 
1. Remove All Adjacent Duplicates In String

```goland
func removeDuplicates(S string) string {
    // 递归
    // 使用栈
    
    stack := make([]byte, 0, len(S))
	for i := range S {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}
```

### 提示
// Use a stack to store the characters, when there are k same characters, delete them.
// To make it more efficient, use a pair to store the value and the count of each character.


### 方法
1. 继承之前的解题方法。
2. 新的思路。