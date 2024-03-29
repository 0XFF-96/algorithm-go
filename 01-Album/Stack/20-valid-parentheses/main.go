package main

// setup 过程
//GOROOT=/Users/lijinrui/go/go1.13.5 #gosetup
//GOPATH=/Users/lijinrui/go #gosetup
///Users/lijinrui/go/go1.13.5/bin/go build -o /private/var/folders/09/r6f6jypj0f9bj7rz2qpv6jyr0000gn/T/___go_build_main_go /Users/lijinrui/algorithm-go/stack/20-valid-parentheses/main.go #gosetup
///private/var/folders/09/r6f6jypj0f9bj7rz2qpv6jyr0000gn/T/___go_build_main_go #gosetup

func main() {
	isValid("s")
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	m := map[byte]byte{
		')':'(',
		'}':'{',
		']':'[',
	}

	// 可以减少类型转换的
	// stack []byte ?
	// 可以减少类型转换的次数。
	var stack []string
	for i:=0;i < len(s);i ++{
		char := s[i]
		v, ok := m[char]
		if ok {
			// we may have edge case when ) was push into stack before any open parenthesis
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if string(v) == top {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		} else {
			// if we have open parentheses pop it in to stack
			stack = append(stack, string(char))
		}
	}
	return len(stack) == 0
}

func helper(a, b string) bool {
	return (a+b == "()") || (a+b == "[]") || (a+b == "{}")
}

func isValidV2(s string) bool {
	str := ""
	for _, t := range s {
		v := string(t)
		if v == "(" || v == "[" || v == "{" {
			str += v
		} else {
			if len(str) == 0 {
				return false
			}
			if !helper(string(str[len(str)-1]), string(v)) {
				return false
			}
			str = str[:len(str)-1]
		}
	}
	return len(str) == 0
}