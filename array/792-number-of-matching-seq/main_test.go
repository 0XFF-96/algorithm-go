package main 
// https://www.youtube.com/watch?v=l8_vcmjQA4g

func numMatchingSubseq(S string, words []string) int {
    // map 记录字母位置。 
    // subsequences , 
    // 不是 continous 。 
    
    // 搜索的一个问题。 
    // 对每个 words 进行一次 dfs 搜索。 
    // N  * O(dfs) 的时间。 
    
    // build a trie ?
    // and 用 S  Match Prefix words lol。 
    // 
    // https://www.youtube.com/watch?v=l8_vcmjQA4g
    // 想法💡差不多，基本能实现的吧。 
}

func numMatchingSubseqV2(S string, words []string) int {
    count := 0
    L := len(S)
    
    wordmap := make(map[string]int)
    for _, word := range words {
        wordmap[word]++
    }
    
    for word, val := range wordmap {
        w := len(word) 
        j := 0
        for i := 0; i < L && j < w; i++ {
            if S[i] == word[j] {
                j++
            } 
            // 不能够 break 
            // 需要继续搜索下去。
        }
    
        if j == w {
            count += val
        }
    }
    return count
}