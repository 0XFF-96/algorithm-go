
// 不太懂，为什么使用 recursive 来做
func isValidSerialization(preorder string) bool {
    list := strings.Split(preorder, ",")
    i:=0
    return validate(list, &i) && i == len(list)
    
}

func validate(preorder []string, i *int) bool {
    if len(preorder) == *i {  //  ？
        return false
    }
    if preorder[*i] == "#" {  // ？
        *i += 1
        return true
	}
	
    *i += 1   
    // i 的明确含义是？
	// 不太可能用在生产环境。
    left := validate(preorder, i)
    right := validate(preorder, i)
    return left && right
}