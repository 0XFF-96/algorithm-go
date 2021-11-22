package trieLeetCode

import (
	"testing"
)

// trie 🌲的一些文章和习题
// https://leetcode.com/tag/trie/



// 题解1：https://leetcode.com/problems/implement-trie-prefix-tree/solution/
// 应用场景：1. Autocomplete 2. Spell checker 3.IP routing (Longest prefix matching)
// 4.

// Another reason why trie outperforms hash table,
// is that as hash table increases in size, there are lots of hash collisions and the search time complexity could deteriorate to O(n)O(n),
// where nn is the number of keys inserted.
// Trie could use less space compared to Hash Table when storing many keys with the same prefix.
// In this case using trie has only O(m)O(m) time complexity, where mm is the key length. Searching for a key in a balanced tree costs O(m \log n)O(mlogn) time complexity.



// solution1:https://leetcode.com/problems/implement-trie-prefix-tree/discuss/338819/go-concise-implement
type Trie struct {
	next map[rune]*Trie
	isWord bool
}

func Constructor() Trie {
	return Trie{next: make(map[rune]*Trie), isWord:false}
}
func (t *Trie) Insert(word string){
	for _, v := range word {
		if t.next[v] == nil {
			t.next[v] = &Trie{next:make(map[rune]*Trie), isWord:false}
		}
		t = t.next[v]
	}
	t.isWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.prefixNode(word)
	return node != nil && node.isWord
}

func (t *Trie) StartWtih(prefix string) bool {
	node := t.prefixNode(prefix)
	return node != nil
}
// 前缀搜索...
//
func (t *Trie) prefixNode(prefix string) *Trie{
	for _, v := range prefix {
		if t.next[v] == nil {

			// 没有这个 keyword...
			//
			return nil
		}
		t = t.next[v]
	}
	return t
}


// Java 风格的写法
// 需要写一大堆的， getChild, setChild, setValue, getValue 等函数。
//
//const trieLen = 26
//
//type Trie struct {
//	children []*Trie
//	value    string
//}


// https://leetcode.com/problems/palindrome-pairs/discuss/158103/Go-trie-O(n*k)-solution-with-explanation
// 336. Palindrome Pairs
//


type node struct {
	child [26]*node
	idx int
	palindromeBranches []int
}

func insert(root *node, word string, idx int){
	cur := root
	for i := len(word)-1;i>=0;i--{
		c := int(word[i] - 'a')
		n := cur.child[c]
		if n == nil {
			n = &node{idx:-1}
			cur.child[c] = n
		}
		cur = n
		if i == 0 {
			cur.idx = idx
		}
	}
}

func buildTrie(words []string) *node {
	root := &node{idx:-1}
	for i, word := range words {
		insert(root, word, i)
	}
	return root
}

func palindrome(word string) bool {
	left := 0
	right := len(word) -1
	for left <= right {
		if word[left] != word[right] {
			return false
		}
		left++
		right --
	}
	return true
}

func branches(prefix string, n *node)([]string, []int){
	var strs []string
	var ints []int

	if n.idx >= 0 {
		strs = append(strs, prefix)
		ints = append(ints, n.idx)
	}

	for i, child := range n.child {
		if child != nil {
			r0, r1 := branches(prefix + string([]byte{byte((i)+'a')}),child)
			strs = append(strs, r0...)
			ints = append(ints, r1...)
		}
	}
	return strs, ints
}

//func palindromeBranches(n *node) []int {
//	if n.palindromeBranches != nil {
//		return n.palindromeBranches
//	}
//
//}





func TestPAC(t *testing.T){
//Input: ["abcd","dcba","lls","s","sssll"]
//Output: [[0,1],[1,0],[3,2],[2,4]]
//Explanation: The palindromes are ["dcbaabcd","abcddcba","slls","llssssll"]

//Input: ["bat","tab","cat"]
//Output: [[0,1],[1,0]]
//Explanation: The palindromes are ["battab","tabbat"]

	res := []string{"abcd", "dcba", "lls", "s", "sssll"}
	r := palindromePairs(res)
	t.Log(r)
}



// go golang.
// solution 2.
// 暴力解法.
// 时间复杂度 N*N*m.
// 优化的方向是第二个 for 循环.
func palindromePairs(words []string) [][]int {
	l := len(words)
	result := make([][]int, 0)
	for i := 0; i < l-1;i++{
		for j := i+1;j <l;j++{
			if isPalindrome(words[i] + words[j]){
				temp := make([]int, 0)
				temp  = append(temp, i)
				temp = append(temp, j)
				result = append(result, temp)
			}
			if isPalindrome(words[j] + words[i]) {
				temp := make([]int, 0)
				temp = append(temp, j)
				temp = append(temp, i)
				result = append(result, temp)
			}
		}
	}

	return result
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s) / 2;i++{
		if s[i] != s[len(s)-1-i]{
			return false
		}
	}
	return true
}



// TODO:Maximum XOR of Two Numbers in an Array
// 每个算法题，加一个一个索引
// 暴力写法, 两个 for 循环

type node2 struct {
	child [2]*node2
}


func insertV2(root *node, n int){
	cur := root
	for i := 0; i <= 30; i++{
		s := uint(30-i)
		c := (n >> s) & 1
		child := cur.child[c]
		if child == nil {
			child = &node{}
			cur.child[c] = child
		}
		cur = child
	}
}

func buildTrieV2(nums []int) *node {
	root := &node{}
	for _, n := range nums {
		insertV2(root, n)
	}
	return root
}


// 为什么是 MaxXOR
func findMaximumXOR(nums []int) int {
	root := buildTrieV2(nums)
	max := 0

	for _, n := range nums {
		val := 0
		cur := root
		for i := 0; i <= 30; i++ {
			s := uint(30 -i)
			c := (n >>s) & 1
			var child *node
			opposite := true

			if c == 0 {
				child = cur.child[1]
				if child == nil {
					child = cur.child[0]
					opposite = false
				}
			} else {
				child = cur.child[0]
				if child == nil {
					child = cur.child[1]
					opposite = false
				}
		 	}

			if child == nil {
				break
			}

			if opposite {
				val = val | ( 1 << s )
			}

			cur = child
		}
		if max < val {
			max = val
		}
	}
	return max

}



// TODO: https://leetcode.com/problems/replace-words/
// https://www.fileformat.info/info/unicode/utf8.htm
// x80 这个常量代表什么？
// 想法💡
// 1、build a trie 树🌲
// 2、用 strings.Feild() 切隔每个单词
// 3、逐个单词与 trie 进行比对, append 与 trie 树搜索出来的结果。
// 4、用最长的进行切割，还是最短的进行切割？

func TestFunc(t *testing.T){
	//t.Log(len("."))
	//s, _ := url.QueryUnescape("%E2%8F%B1%20")
	//t.Log(s)

	t.Log(len("好的时候呢，17页的时候呢，就收我们把我们前面这个新功能"))

	t.Log(len("时空的"))
}