package tree

// 要具有框架思想看问题
// https://fudonglai.github.io/2018/12/22/%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E8%87%AA%E9%A1%B6%E5%90%91%E4%B8%8B%E6%A6%82%E8%A7%88/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {return true }
	if p == nil || q == nil || p.Val != q.Val { return false}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 我的解法
func isSameTreeV2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right) &&
		p.Val == q.Val
}


// 思考题：如何使用非递归去做呢？
// 1、需要用到栈 2、压栈 弹出
// https://leetcode.com/problems/same-tree/discuss/327852/Golang-Recursive-and-BFS-solutions
// 非递归版本
func isSameTreeWithStack(p *TreeNode, q *TreeNode) bool {
	pstack := []*TreeNode{p}
	pcur := p
	//qstack := []*TreeNode{q}
	//qcur := q

	for len(pstack) != 0 || pcur != nil {
		if pcur == nil {
			pcur = pstack[len(pstack)-1]
			pstack = pstack[:len(pstack)-1]
		}
		if pcur.Left != nil {
			pstack = append(pstack, pcur)
			pcur = pcur.Left
		}
	}

	return false
}


// 这种做法清晰
// 当时浪费存储空间
func isSameTreeWithQueue(p, q *TreeNode) bool {
	var deq [][]*TreeNode
	deq = append(deq, []*TreeNode{p, q})
	for len(deq) != 0 {
		queue := deq[len(deq)-1]
		deq = deq[:len(deq)-1]
		p = queue[0]
		q = queue[1]

		if !check(p, q) {
			return false
		}
		if p != nil {
			deq = append(deq, []*TreeNode{p.Left, q.Left})
			deq = append(deq, []*TreeNode{p.Right, q.Right})
		}

	}
	return true
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil )||
		(p != nil && q == nil ) {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return true
}

// 利用前序遍历的写法就够了
// https://leetcode-cn.com/problems/same-tree/solution/golangdie-dai-jie-fa-by-fir/


