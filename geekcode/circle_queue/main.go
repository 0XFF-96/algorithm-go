package circle_queue


// 解决链表相关的题目
// 1、警惕指针丢失和内存泄漏(单链表)
// 2、警惕指针丢失和内存泄漏(单链表)
// 相关提醒.⏰
// 1.如果链表为空时，代码是否能正常工作?
// 2.如果链表只包含一个节点时，代码是否能正常工作?
// 3.如果链表只包含两个节点时，代码是否能正常工作?
// 4.代码逻辑在处理头尾节点时是否能正常工作?

// 常见题目
// 1.单链表反转 2.链表中环的检测 3.两个有序链表合并 4.删除链表倒数第n个节点 5.求链表的中间节点
// 2.

// 用数组实现的循环♻️队列
type CircularQueue struct {
	item []string
	capacity int
	head int
	tail int
}

func NewCircularQueue(capacity int) *CircularQueue{
	return &CircularQueue{
		item: make([]string, capacity),
		capacity:    capacity,
		head: 0,
		tail: 0,
	}
}

func (c *CircularQueue) enqueue(item string) bool {
	// 判断队列满的条件
	if ((c.tail + 1) % c.capacity )== c.head {
		return false
	}
	c.item[c.tail]= item
	c.tail = (c.tail + 1) % c.capacity
	return true
}

func (c *CircularQueue) dequeue(item string) string  {
	// 如果 head == tail 表示队列为空🈳️
	if (c.head == c.tail) {
		// 需要定一个常量
		return ""
	}
	ret := c.item[c.head]
	c.head = (c.head +1) % c.capacity
	return ret
}
