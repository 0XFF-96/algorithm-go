
package main 

func helperReverse(s []byte) string {
	for i:=0;i< len(s)/2;i++{
		s[i],s[len(s)-1-i] = s[len(s)-1-i],s[i]
	}
	return string(s)
}

func reverseParentheses(s string) string {
	var res string
	stack :=[]string{}
	for _,c := range s{
		if c == '('{
			stack = append(stack,res)
			res =""
		}else if c== ')'{
			res = helperReverse([]byte(res))
			res = stack[len(stack)-1] + res
			stack = stack[0 :len(stack)-1]
		}else {
			res = res + string(c)
		}
	}
	return res
}