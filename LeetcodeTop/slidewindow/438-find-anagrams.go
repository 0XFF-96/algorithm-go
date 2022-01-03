

// #438 #findAnagrams
// 找到字符串中所有字母异词


// 具体的，我们可以先创建一个大小为 2626 的数组 c2c2 来统计字符串 p 的词频，另外一个同等大小的数组 c1c1 用来统计「滑动窗口」内的 s 的子串词频。

// 当两个数组所统计词频相等，说明找到了一个异位组，将窗口的左端点加入答案

// 核心
// right当前遍历到的字符加入s_cnt后不满足p_cnt的字符数量要求

// 思想
// 1. 字母异词，字母的 count 是一样的。 子数组中，字母出现的 count 是一样的。 
// 2. 子串，转化为计算 count 。 比较难想。 

// 变长窗口版本。diff也可以不预先统计，将它转变为遍历s的时候一种“消耗品”——当“消耗品”不足，
// 唯一可以做的就是移动左窗口释放一些出来。
// 这里面有一个比较特殊的corner case，s中有一个p中没有的字符。下面的代码刚好自洽：


func findAnagrams(s, p string) (ans []int) {
    sLen, pLen := len(s), len(p)
    if sLen < pLen {
        return
    }

    var sCount, pCount [26]int
    for i, ch := range p {
        sCount[s[i]-'a']++
        pCount[ch-'a']++
    }
    if sCount == pCount {
        ans = append(ans, 0)
    }

    for i, ch := range s[:sLen-pLen] {
        sCount[ch-'a']--
        sCount[s[i+pLen]-'a']++
        if sCount == pCount {
            ans = append(ans, i+1)
        }
    }
    return
}