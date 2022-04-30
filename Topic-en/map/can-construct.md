# Can Construct 

Given two strings ransomNote and magazine, return true if ransomNote can be constructed from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.

```

// func canConstruct(ransomNote string, magazine string) bool {
//     // 这道题目和242.有效的字母异位词很像，242.有效的字母异位词相当于求 字符串a 和 字符串b 是否可以相互组成 ，
//     // 而这道题目是求 字符串a能否组成字符串b，而不用管字符串b 能不能组成字符串

//     // 为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，
//     // 组成单词来表达意思”  这里*说明杂志里面的字母不可重复使用。*
//     // 你可以假设两个字符串均只含有小写字母。” *说明只有小写字母*，这一点很重要
// }


func canConstruct(ransomNote, magazine string) bool {
    if len(ransomNote) > len(magazine) {
        return false 
    }

    // 代替哈希表，节省空间
    cnt := [26]int{}

    for _, ch := range magazine {
        cnt[ch - 'a']++
    }

    for _, ch := range ransomNote {
        cnt[ch - 'a']--
        if cnt[ch - 'a'] < 0 {
            return false 
        }
    }
    return true 
}

```
