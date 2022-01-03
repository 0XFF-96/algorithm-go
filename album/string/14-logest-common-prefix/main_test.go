package main 

// vertical solution
func longestCommonPrefix(strs []string) string {
    if len(strs) < 1 {
        return ""
    }
    
    // input: "aaab", "aab", "abc" output:
    // input: "aaaa", "aaa", "aaaab" output: "aaa"
    var prefix string
    str := strs[0]
    for i := 0; i < len(str);i++ {
        prefix = string(str[:i+1])
        for _, s := range strs {
            if !strings.HasPrefix(s, prefix) {
                return string(prefix[:len(prefix)-1])
            }
        }
    }
    return prefix
}


// horizantoal solution ?