package trieLeetCode

import (
	"math"
	"sort"
)

// TODO: Longest Word in Dictionary
// lexicographically 字典排序
//  If there is more than one possible answer,
//  return the longest word with the smallest lexicographical order.
// 暴力方法？1、先进行排序 2、
// 🈶
func longestWord(words []string) string {
	if len(words) < 2 { return ""}

	tmp := []string{}
	max := math.MinInt32

	// 排序
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	// 从最短的开始
	for i := len(words) -1; i > 0; i -- {
		count := 1
		l := len(words[i])
		for j := i -1; j >=0; j -- {
			if words[i][:l-count] == words[j]{
				count ++
			}
			if count == l {}
			if len(words[i]) > max {
				max = len(words[i])
			}
			if len(words[i]) == max {
				tmp = append(tmp, words[i])
			}
		}
	}
	if len(tmp) == 1 { return tmp[0]}
	res := tmp[0]
	for _, v := range tmp {
		if v < res {
			res = v
		}
	}
	return res
}

// TODO: Concatenated Words
// 拼接词
// https://leetcode.com/problems/concatenated-words/
// https://leetcode.com/problems/concatenated-words/discuss/349281/Simple-recursive-GO-solution-that-is-faster-than-96.05-and-use-less-memory-than-100.00
func findAllConcatenatedWordsInADict(words []string) []string {
	candiateMap := make(map[string]bool)
	res := make([]string, 0)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	for _, word := range words {
		if verify(word, candiateMap) {
			res = append(res, word)
		}
		candiateMap[word] = true
	}
	return res
}

func verify(word string, candiateMap map[string]bool) bool {
	if candiateMap[word] {
		return true
	}

	for i := 1; i < len(word); i++{
		pre := word[:i]
		suffix := word[i:]
		if candiateMap[pre]{
			if candiateMap[suffix]{
				return true
			}
			valid := verify(suffix, candiateMap)
			if valid {
				candiateMap[suffix] = true
				return true
			}
		}
	}
	return false
}







// TODO: overlap
// https://medium.com/algorithm-and-datastructure/index-pairs-of-a-string-7b7c8306ead0






