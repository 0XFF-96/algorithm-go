package array



// 2、有两种方法。https://leetcode.com/problems/find-common-characters/discuss/261020/Golang-solution
// 1、https://leetcode.com/problems/find-common-characters/discuss/250429/golang-recursive-method
// 递归方法。
//

// 错误的想法。
//
func commonChars(A []string) []string {
	// map count letter
	// len(A)
	m := make(map[string]int)

	for _, v := range A {
		for i:=0;i<len(v);i++{
			m[string(v[i])]++
		}
	}

	la := len(A)
	var ret []string
	for k, v := range m {
		tmp := v/la
		if tmp >= 1 {
			for i:=tmp;i >= 1;i--{
				ret = append(ret, k)
			}
		}
	}

	return ret
}


