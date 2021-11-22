package trieLeetCode

import (
	"container/heap"
	"testing"
)

type Item struct {
	value string
	priority int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {return len(pq)}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].priority == pq[j].priority {
		return pq[i].value < pq[j].value
	}
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop () interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	// avoid memory leak
	old[n-1] = nil
	*pq = old[0:n-1]
	return item
}

func topKFriquent(words []string, k int) []string {
	frequentMap := map[string]int{}

	lenUniqueWords := 0
	for _, word := range words {
		if frequentMap[word] == 0 {
			// 为什么要知道这个信息？
			lenUniqueWords++
		}
	frequentMap[word]++
	}

	pq := make(PriorityQueue, lenUniqueWords)
	i := 0
	for value, priority := range frequentMap {
		pq[i]=&Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}

	heap.Init(&pq)
	var resp []string

	for j := 0; j < k; j ++ {
		item := heap.Pop(&pq).(*Item)
		resp = append(resp, item.value)
	}
	return resp
}


// TODO: Top K Frequent Words
func TestTopK(t *testing.T){
	words := []string{"123", "123", "3", "1", "1", "1"}
	w := topKFriquent(words, 3)
	t.Log(w)
}


// https://leetcode.com/problems/top-k-frequent-words/discuss/493291/golang-solution-using-heap-and-hashmap
// 另一种实现方式

