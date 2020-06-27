// "3[a]2[bc]"
// this is the test case. 
// 
func decodeString(s string) string {
	counts := []int{}
	strings := []string{}
	index := 0
	res := ""
	for index < len(s) {
		if s[index] >= '0' && s[index] <= '9' {
			count := 0
			for s[index] >= '0' && s[index] <= '9' {
				count = count*10 + int(s[index]-'0')
				index++
			}
			counts = append(counts, count)
		} else if s[index] == '[' {
			strings = append(strings, res)
			res = ""
			index++
		} else if s[index] == ']' {
			n := counts[len(counts)-1]
			counts = counts[:len(counts)-1]
			tmp := strings[len(strings)-1]
			strings = strings[:len(strings)-1]
			for i := 0; i < n; i++ {
				tmp += res
			}
			res = tmp
			index++
		} else {
			res += string(s[index])
			index++
		}
	}
	return res
}