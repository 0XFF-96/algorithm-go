package mian 

import (
    "strings"
)

func minRemoveToMakeValid(s string) string {
	pairCount := getPairCount(s)
	open, closee := pairCount, pairCount

	sb := strings.Builder{}
	for _, r := range s {
		if r == '(' && open > 0 {
            open--
			sb.WriteRune('(')
		} else if r == ')' && closee > open {
            closee--
			sb.WriteRune(')')
		} else if r != '(' && r != ')' {
			sb.WriteRune(r)
		}
	}
    
	return sb.String()
}

func getPairCount(s string) int {
	open, closee := 0, 0

	for _, r := range s {
		if r == '(' {
			open++
		} else if r == ')' && closee < open {
			closee++
		}
	}

	if open < closee {
		return open
	}
	return closee
}

func minRemoveToMakeValidV2(s string) string {
    var st []int  //save parentheses index which should be delete
	var res string = ""
    for i, ch := range s {
        if ch == '(' {
            st = append(st, i)
        }else if ch == ')' {
            if len(st) > 0 && s[st[len(st) - 1]] == '(' {
                st = st[:len(st) - 1]
            }else {  //stack empty or top is ')'
                st = append(st, i)
            }
        }
    }
    start, end := 0, len(s)
    for len(st) > 0 {
        start = st[len(st) - 1] + 1
        st = st[:len(st) - 1]  
        res = s[start:end] + res     
        end = start - 1
        start = 0
    }
    res = s[start:end] + res
    return res
}