package tree

// https://www.youtube.com/watch?v=LR3K5XAWV5k
// 视频资料
//
type Record struct {
	first, second, prev *TreeNode
}

// 递归实现
func inorder(t *TreeNode, r *Record){
	if t == nil {
		return
	}

	inorder(t.Left, r)
	if r.prev != nil && r.prev.Val > t.Val {
		if r.first == nil{
			r.first, r.second = r.prev, t
		} else {
			r.second = t
		}
	}
	r.prev = t

	inorder(t.Right, r)
}

func recoverTree(root *TreeNode){
	r := &Record{}
	inorder(root, r)

	if r.first != nil {
		r.first.Val, r.second.Val = r.second.Val, r.first.Val
	}
}


func recoverTreeWithStack(root *TreeNode){
	var  pre, firstStartNode, _ *TreeNode
	cur := root
	// var stack []*TreeNode
	stack := []*TreeNode{cur}
	for len(stack) != 0 || cur != nil {
		if cur == nil {
			tmp := stack[len(stack)-1]
			stack = stack [:len(stack)-1]
			if pre != nil && pre.Val > tmp.Val {
				if firstStartNode == nil {
					firstStartNode = pre
				}
				_ = tmp
			}
			pre = tmp
			if tmp.Right != nil {
				cur = tmp.Right
			}
			continue
		}

		stack = append(stack, cur)
		cur = cur.Left
	}
	// 最后交换两个数据的指针...
}



// 这个实现有点问题
// 有没有非递归的实现？
func recoverTreeV2(root *TreeNode){
	stack := make([]*TreeNode, 0)
	var prev *TreeNode
	var firstStartPointer *TreeNode
	var lastEndPointer *TreeNode
	cur := root

	for len(stack)!= 0|| cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if prev != nil && prev.Val > temp.Val {
				if firstStartPointer == nil {
					firstStartPointer = prev
				}
				lastEndPointer = temp
			}
			prev = temp

			if temp.Right != nil {
				cur = temp.Right
			}
		}
	}
	firstStartPointer.Val,lastEndPointer.Val = lastEndPointer.Val, firstStartPointer.Val
}







