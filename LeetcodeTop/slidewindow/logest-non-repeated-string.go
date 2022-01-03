 
 
 // 无重复字符的最长子串
 // 

 // 1. hast set ， value 记录的是出现次数，还是记录的是 index 下标值。 
 // 2. 滑动窗口。

 // 变量名字可以起好一点！
 // 
 func lengthOfLongestSubstring(s string) int {
    // 哈希集合，记录每个字符是否出现过
    m := map[byte]int{}  // 这个哈希集合设置得很巧妙, int 记录的是 index 的值。下标位置。

    n := len(s)
    // 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
    rk, ans := -1, 0
    for i := 0; i < n; i++ {
        if i != 0 {
            // 左指针向右移动一格，移除一个字符
			// 这里是冗余的!
            delete(m, s[i-1])
        }
        for rk + 1 < n && m[s[rk+1]] == 0 {
            // 不断地移动右指针
            m[s[rk+1]]++
            rk++
        }
        // 第 i 到 rk 个字符是一个极长的无重复字符子串
        ans = max(ans, rk - i + 1)
    }
    return ans
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

```更好的解法
class Solution {
    public int lengthOfLongestSubstring(String s) {
        if (s.length()==0) return 0;
        HashMap<Character, Integer> map = new HashMap<Character, Integer>();
        int max = 0;
        int left = 0;
        for(int i = 0; i < s.length(); i ++){
            if(map.containsKey(s.charAt(i))){
                left = Math.max(left,map.get(s.charAt(i)) + 1);
            }
            map.put(s.charAt(i),i);
            max = Math.max(max,i-left+1);
        }
        return max;
        
    }
}
```

// 缺点， 答案有个缺点，左指针并不需要依次递增，即多了很多无谓的循环。 发现有重复字符时，可以直接把左指针移动到第一个重复字符的下一个位置即可。
// 每次左指针右移一位，移除set的一个字符，这一步会导致很多无用的循环。while循环发现的重复字符不一定就是Set最早添加那个，还要好多次循环才能到达，这些都是无效循环，不如直接用map记下每个字符的索引，直接进行跳转。 
