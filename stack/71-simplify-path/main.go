package _1_simplify_path

import (
	"strings"
)

func simplifyPath(path string) string {
	// canonical path
	// absolute path
	// pop and push 的栈操作
	paths := strings.Split(path, "/")
	var stack []string

	for _, str := range paths {
		switch str {
		case ".":
		case "..":
			// pop
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case "":
		default:
			stack = append(stack, str)
		}
	}
	// fmt.Println(stack)
	if len(stack) == 0 {
		return "/"
	}

	var ret string
	for _, str := range stack {
		ret += "/" + str
	}
	return ret
}
