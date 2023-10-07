package tree


// https://leetcode.com/problems/check-completeness-of-a-binary-tree/solution/
// 不会
// 和树的 position 的那一题非常像
//
// This is the scheme we will use. Under this scheme, our tree is complete if the codes take the form 1, 2, 3, ... in sequence with no gaps
//
//



// 检查所有节点，
// 如果只有左节点，而没有右节点则..
func isCompleteTree(root *TreeNode) bool {

	aNode := &Anode{node:root, pos:1}
	queue := []*Anode{aNode}

	seq := []*Anode{}

	// bfs
	for len(queue) != 0 {
		size := len(queue)
		for i:=0;i<size;i++{
			top := queue[0]
			queue = queue[1:]

			if top.node.Left != nil {
				queue = append(queue, &Anode{node:top.node.Left, pos: top.pos * 2})
			}

			if top.node.Right != nil {
				queue = append(queue, &Anode{node:top.node.Right, pos: top.pos *2 +1})
			}

			seq = append(seq, top)
			// 可以在这里做一个判断
		}
	}

	for i:=1;i<len(seq);i++ {
		if seq[i].pos - seq[i-1].pos > 1 {
			return false
		}
	}
	return true

}

type Anode struct {
	node *TreeNode
	pos int
}

// 优化的想法
// 1、优化最后一次 for 循环，将其移动到一直遍历的 for 中
// 2、seq 的空间使用，将其改为两个变量的值。
//

