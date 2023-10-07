package main

type item struct {
    next, count int
}

type pq []*item
func (p pq) Len() int { return len(p) }
func (p pq) Less(i, j int) bool { 
    if p[i].next != p[j].next {
        return p[i].next < p[j].next
    }
    return p[i].count >= p[j].count
}
func (p pq) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p *pq) Push(x interface{}) { *p = append(*p, x.(*item)) }
func (p *pq) Pop() interface{} {
    x := (*p)[len(*p) - 1]
    *p = (*p)[0:len(*p) - 1]
    return x
}

func leastInterval(tasks []byte, n int) int {
    var pq pq
    m := make(map[byte]int)
    for _, v := range tasks {
        m[v]++
    }
    for _, v := range m {
        item := &item{0, v}
        pq = append(pq, item)
    }
    heap.Init(&pq)
    sz := -1
    for len(pq) > 0 {
        sz++
        x := heap.Pop(&pq).(*item)
        if x.next > sz {
            sz += (x.next - sz)
        }
        x.count--
        if x.count > 0 {
            x.next += n + 1
            heap.Push(&pq, x)
        }
    }
    return sz + 1
}