package tree


// 层次遍历的最后一个节点
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var ret []int
	for len(queue) > 0 {
		// 重点和难点都在于这里都边界条件
		// if i == size
		size := len(queue)-1
		for i := 0; i <= size;i++{
			// dequeue
			top := queue[0]
			queue = queue[1:]
			if i == size {
				ret = append(ret,top.Val)
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}

		}

	}
	return ret
}



// 不太了解非递归解法的特别之处...

func rightSideView2(root *TreeNode) []int {
	result:= make([]int, 0)
	rightView2(root, &result, 0)
	return result
}


func rightView2(curr *TreeNode, result *[]int, currDepth int) {
	if curr == nil {
		return
	}

	// 这个判断的依据是什么？
	// 我知道了
	// 从右边开始遍历起来...
	// 当有结果了，就越界，然后就不会加入到队列里面了
	// 真他娘的是个人才...
	// 鬼才
	if currDepth == len(*result) {
		*result = append(*result, curr.Val)
	}

	rightView2(curr.Right, result, currDepth + 1 )
	rightView2(curr.Left, result, currDepth + 1)
}

// 层次遍历的递归方式是什么？
// 有什么好处呢？
//