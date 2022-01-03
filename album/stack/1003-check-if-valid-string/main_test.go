package main

// Input: "aabcbc"
// Output: true
// Input: "abcabcababcc"
// Output: true
// Input: "abccba"
// Output: false

func isValid(S string) bool {
    stack := []rune{}
        
    for _, s := range S {
        if s == 'c' {
            if len(stack) == 0 {
                return false
            }
            pop := stack[len(stack)-1]
            if pop != 'b' {
                return false
            }
            stack = stack[:len(stack)-1]
            
            if len(stack) == 0 {
                return false
            }
            pop = stack[len(stack)-1]
            if pop != 'a' {
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s)
        }
    }
    return len(stack) == 0 
}


    // public boolean isValid(String s) {
    //     Stack<Character> stack = new Stack<>();
    //     for (char c: s.toCharArray()) {
    //         if (c == 'c') {
    //             if (stack.isEmpty() || stack.pop() != 'b') return false;
    //             if (stack.isEmpty() || stack.pop() != 'a') return false;
    //         } else {
    //             stack.push(c);
    //         }
    //     }
    //     return stack.isEmpty();
    // }
