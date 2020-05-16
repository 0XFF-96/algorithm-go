package main

// 利用 map 维护原 node 和 新 node 之间的关系
// O(n) 和 O(n) 的方法
func copyRandomListV3(head *Node) *Node {
	cur := head
	dummy := &Node{}
	dCur := dummy
	visited := map[*Node]*Node{}
	// copy list first
	for cur != nil {
		n := &Node{Val: cur.Val}
		visited[cur] = n
		dCur.Next = n
		dCur = dCur.Next
		cur = cur.Next
	}
	// copy random pointers
	cur, dCur = head, dummy
	for cur != nil {
		dCur.Next.Random = visited[cur.Random]
		dCur = dCur.Next
		cur = cur.Next
	}
	return dummy.Next
}
