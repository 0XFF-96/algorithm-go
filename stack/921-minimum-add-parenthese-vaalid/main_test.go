

func minAddToMakeValid(S string) int {
    // parentheses valid 
    // minimum to add
    ans := 0 
    bal := 0 
    for _, s := range S {
        if string(s) == "("{
            bal += 1
        } else {
            bal -= 1
        }
		// it is guarantedd bal >= -1
		// 
        if bal == -1 {
            ans +=1 
            bal +=1
        }
    }
    return ans + bal 
}