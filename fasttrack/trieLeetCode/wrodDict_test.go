package trieLeetCode

import (
	"container/list"
	"testing"
)

// https://leetcode.com/problems/add-and-search-word-data-structure-design/
// word 的一些设计

//void addWord(word)
//bool search(word)

// Trie + DFS
// how to implement . that can represent any characters?
// ...

type WordDictionary struct {
	Children [26]*WordDictionary
	IsWord bool
}

func ConstructorW() WordDictionary{
	return WordDictionary{Children:[26]*WordDictionary{}, IsWord:false}
}

func (this *WordDictionary) AddWord(word string){
	node := this
	for _, char := range word {
		if node.Children[char - 'a'] == nil {
			node.Children[char - 'a'] = new(WordDictionary)
		}
		node = node.Children[char - 'a']
	}
	node.IsWord = true
}

// Returns if the word is in the data structure
// A word could contain the dot character '.'
// to represent any one letter

func (this *WordDictionary) Search(word string) bool {
	var (
		queue []*WordDictionary
		indices []int
		node *WordDictionary
		i int
	)
	queue = append(queue, this)
	indices = append(indices, 0)
	for len(queue) > 0 {
		// 这两行用来干嘛？
		node, queue = queue[0], queue[1:]
		i, indices = indices[0], indices[1:]

		// 退出条件
		if i == len(word){
			if node.IsWord{
				return true
			}
			continue
		}
		// DFS
		if rune(word[i]) == '.'{
			for _ ,child := range node.Children {
				if child != nil {
					queue = append(queue, child)
					indices = append(indices, i + 1)
				}
			}
		} else if node.Children[rune(word[i]) - 'a'] != nil {
			queue = append(queue, node.Children[rune(word[i]) - 'a'])
			indices = append(indices, i + 1)
		}
	}
	return false
}

func TestWordSearch(t *testing.T){
	// word search
}


func (this *WordDictionary) SearchViaList(word string) bool {
	queue, indices := list.New(), list.New()
	queue.PushBack(this)
	indices.PushBack(0)

	// #TODO: 为什么只用判断 e1 != nil ?
	for e1, e2 := queue.Front(), indices.Front(); e1 != nil; e1, e2 = e1.Next(), e2.Next(){
		node, i := e1.Value.(*WordDictionary), e2.Value.(int)


		// DFS 的退出条件
		if i == len(word){
			if node.IsWord{
				return true
			}
			continue
		}
		if rune(word[i]) == '.' {
			for _, child := range node.Children{
				if child != nil {
					queue.PushBack(child)
					indices.PushBack(i+1)
				}
			}
		} else if node.Children[rune(word[i]) - 'a'] != nil {
			queue.PushBack(node.Children[rune(word[i]) - 'a'])
			indices.PushBack(i + 1)
		}
	}
	return false
}



// Trie 树的性质
// 以及相关习题
// https://wizardforcel.gitbooks.io/the-art-of-programming-by-july/content/06.09.html