package tree


// https://leetcode.com/problems/complete-binary-tree-inserter/
//

type CBTInserter struct {
	treeArr []*TreeNode
}

func ConstructorV2(root *TreeNode) CBTInserter {
	return CBTInserter{treeArr: bsf(root)}
}

func bsf(root *TreeNode) []*TreeNode{
	var queue []*TreeNode
	var resp  []*TreeNode
	queue = append(queue, root)
	for {
		node := queue[0]
		queue = queue[1:]
		resp = append(resp, node)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if len(queue) == 0  {
			break
		}
	}
	return resp
}

func (this *CBTInserter) Insert(v int) int {
	length := len(this.treeArr)
	insert_loc := ((float64(length) - 1) / 2.0 )
	loc := int64(insert_loc)
	parent := this.treeArr[loc]

	node := &TreeNode{Val: v}
	if parent.Left == nil {
		parent.Left = node
	} else {
		parent.Right = node
	}
	this.treeArr = append(this.treeArr, node)
	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.treeArr[0]
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */