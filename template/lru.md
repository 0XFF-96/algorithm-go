### LRU template 

### LRU 的实际应用场景？
- Redis go 里面的 链接池、Go Client 里面的链接池

### LRU 和 LFU 应用场景

LRU和LFU都是内存管理的页面置换算法。

LRU，即：最近最少使用淘汰算法（Least Recently Used）。LRU是淘汰最长时间没有被使用的页面。
LFU，即：最不经常使用淘汰算法（Least Frequently Used）。LFU是淘汰一段时间内，使用次数最少的页面。

LRU算法适合：较大的文件比如游戏客户端（最近加载的地图文件） LFU算法适合：较小的文件和教零碎的文件比如系统文件、应用程序文件 其中：LRU消耗CPU资源较少，LFU消耗CPU资源较多。2、LFU作为负载均衡算法：保证每次使用都是最不经常使用的节点。 1、LFU作为缓存算法。 
https://halfrost.com/lru_lfu_interview/ 。 

### 自然语言描述

双链表中实现的方法：
	链表初始化
	添加元素到双链表尾部（同时意味着该元素最近使用过）
	删除某个结点（非头结点）
	删除并返回头结点（意味着移除最久未使用过的元素）
	返回链表当前长度
LRU 缓存中的方法
	初始化，get，put方法
	设置某元素最近已使用makeRecently（原map中已有该元素）
	添加最近使用过的元素addRecently（原map中不存在该键值对，新元素）
	删除某个key对应的元素
	删除最久未使用过的元素 2.中的方法也可以直接在get和put方法中实现，可以减少部分冗余


### 简短版 LRU 

type LRUCache struct {
    size int
    capacity int
    cache map[int]*DLinkedNode
    head, tail *DLinkedNode
}

type DLinkedNode struct {
    key, value int
    prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
    return &DLinkedNode{
        key: key,
        value: value,
    }
}

func Constructor(capacity int) LRUCache {
    l := LRUCache{
        cache: map[int]*DLinkedNode{},
        head: initDLinkedNode(0, 0),
        tail: initDLinkedNode(0, 0),
        capacity: capacity,
    }
    l.head.next = l.tail
    l.tail.prev = l.head
    return l
}


### From golang-snippet 

	keyLru struct {
		// if would be better if we use capacity as name
		limit    int
		evicts   *list.List
		elements map[string]*list.Element

		// 钩子函数, 用于在
		onEvict  func(key string)

		// 需要并发安全的功能时候，需要加锁
		// 读写锁🔒， 再高级一点，需要并发量大的， 【分段锁】

		// 再再硬核一点的，就需要 【无锁并发安全】
		// 需要绑定系统线程，CPU 的一些 hard 代码。
	}

// 主要是 list.ELement 的双链表实现
// 

func (klru *keyLru) add(key string) {
	if elem, ok := klru.elements[key]; ok {
		klru.evicts.MoveToFront(elem)
		return
	}

	// Add new item
	elem := klru.evicts.PushFront(key)
	klru.elements[key] = elem

	// Verify size not exceeded
	if klru.evicts.Len() > klru.limit {
		klru.removeOldest()
	}
}

func (klru *keyLru) remove(key string) {
	if elem, ok := klru.elements[key]; ok {
		klru.removeElement(elem)
	}
}

func (klru *keyLru) removeOldest() {
	elem := klru.evicts.Back()
	if elem != nil {
		klru.removeElement(elem)
	}
}

func (klru *keyLru) removeElement(e *list.Element) {
	klru.evicts.Remove(e)
	key := e.Value.(string)
	delete(klru.elements, key)
	klru.onEvict(key)
}


// list.Element

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}

// Element is an element of a linked list.
type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element

	// The list to which this element belongs.
	list *List

	// The value stored with this element.
	Value interface{}
}



