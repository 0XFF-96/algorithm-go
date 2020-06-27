package removeOutmostParentheses

import (
	"fmt"
	"strings"
)

// 错误的做法
func removeOuterParenthesesW(S string) string {
	fmt.Println([]rune(S))

	stack := []rune{}
	m := map[int]struct{}{}

	for idx, s := range []rune(S) {
		if len(stack) == 0 {
			m[idx] = struct{}{}
			stack = append(stack, s)
			continue
		}

		if stack[len(stack)-1] == '(' && s == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				m[idx] = struct{}{}
			}
		}

		if stack[len(stack)-1] == ')' && s == '(' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				m[idx] = struct{}{}
			}
		} else {
			stack = append(stack, s)
		}
	}


	// 然后再遍历一遍 S

	return S
}

// case1
// Input: "(()())(())"
// Output: "()()()"

// case2
// Input: "()()"
// Output: ""
// 太牛 B 的做法了
// 简洁，高效
func removeOuterParentheses(S string) string {
	result := ""
	counter := 0

	// 开始位置很重要?
	// 为什从 1 开始？, 而不是从零？
	begin := 1
	for i, s := range S {
		if s == '(' {
			counter++
		} else {
			counter--
		}
		if counter == 0 {
			result += S[begin:i]
			begin = i + 2
		}
	}
	return result
}



func removeOuterParenthesesV2(S string) string {
	stack := ""
	openParan := 0
	var res string
	for _, char := range S {
		c := string(char)
		if c == "(" {
			openParan++
			stack += c
		} else if c == ")" {
			openParan--
			stack += c
		}
		if openParan == 0 {
			// 隐含条件 stack 只有两个元素时，不会加入
			// stack['(', ')']
			// '' = stack[1:len(stack)-1]
			// 这个点，学到了
			// 有点 hack
			// 为什么从 1 开始？
			res += stack[1:len(stack)-1]
			stack = ""
		}
	}
	return res
}



// 有点厉害了
// 1、怎么办？
func removeOuterParenthesesV3(S string) string {
	b := strings.Builder{}
	depth := 0

	for _, ch := range S {
		if ch == '(' {
			if depth != 0 {
				b.WriteRune(ch)
			}
			depth++
		} else {
			if depth != 1 {
				b.WriteRune(ch)
			}
			depth--
		}
	}

	return b.String()
}