package tree


func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := &TreeNode{Val:pre[0]}
	if len(pre) ==1 {
		return root
	}

	// 如果不是满二叉树怎么办？
	// find the
	l := 0
	for i:= 0;i <len(pre) ;i++{
		if post[i] == pre[1] {
			l = i+1
		}
	}

	root.Left = constructFromPrePost(pre[1:l+1], post[0:l])
	root.Right = constructFromPrePost(pre[l+1:], post[l:len(post)-1])

	return root
}
