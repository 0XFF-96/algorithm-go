import "strings"

func isValidSerialization(preorder string) bool {
	strings := strings.Split(preorder, ",")
	stack := []string{}
	for _, str := range strings {
		if str != "#" {
			stack = append(stack, str)
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] == "#" {
			if len(stack) < 2 {
				return false
			}

			stack = stack[:len(stack)-2]
		}

		stack = append(stack, "#")
	}

	return len(stack) == 1 && stack[0] == "#"
}