package tree


// 这是一个 树结构的类型题目
// 操作型的题目
//
//

// 有一次出现第几失误
// 看错了错了语句
// 判断条件也写错了
//
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil || root2 != nil {
		return false
	}

	leftFilp  := flipEquiv(root1.Left, root2.Right)
	left      := flipEquiv(root1.Left, root2.Left)

	right := flipEquiv(root1.Right, root2.Right)
	rightFilp := flipEquiv(root1.Right, root2.Left)

	if root1.Val != root2.Val {
		return false
	}

	// 这个判断语句有问题
	//
	return (left || leftFilp) && ( right || rightFilp)
}


func flipEquivV2(root1 *TreeNode, root2 *TreeNode) bool {
	//     if root1 == nil && root2 == nil {
	//         return true
	//     }
	//     if root1 != nil || root2 != nil {
	//         return false
	//     }

	//     if (root1 == nil || root2 == nil) || root1.Val != root2.Val {
	//         return false
	//     }


	//     left      := flipEquiv(root1.Left, root2.Left)
	//     right      := flipEquiv(root1.Right, root2.Right)

	//     leftF := flipEquiv(root1.Left, root2.Right)
	//     rightF := flipEquiv(root2.Right, root2.Left)


	//     if root1.Val != root2.Val {
	//         return false
	//     }

	if  root1== nil && root2 == nil {
		return true
	}
	if (nil == root1 || nil == root2) || (root1.Val != root2.Val){
		return false
	}

	left := flipEquiv(root1.Left,root2.Left)
	right := flipEquiv(root1.Right,root2.Right)

	leftF := flipEquiv(root1.Left,root2.Right)
	rightF := flipEquiv(root1.Left,root2.Right)

	return (left && right)  || (leftF && rightF)
	// return (flipEquiv(root1.Left,root2.Left) && flipEquiv(root1.Right,root2.Right)) || (flipEquiv(root1.Left,root2.Right) && flipEquiv(root1.Right,root2.Left))

	// 这里面的判断条件
	//
	// return (left && right) || (leftF && rightF)
}
