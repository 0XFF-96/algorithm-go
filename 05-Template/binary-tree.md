### 二叉树

### 重点
1、递归和迭代，二叉树的遍历写法
2、

### 二叉树的遍历框架

- 递归和迭代的解法都需要能够写得出来。 

/* 二叉树遍历框架 */
void traverse(TreeNode root) {
    // 前序遍历
    traverse(root.left)
    // 中序遍历
    traverse(root.right)
    // 后序遍历
}


### 前序、中序、后序遍历的本质

1、前序遍历的代码在进入某一个节点之前的那个时间点执行，
后序遍历代码在离开某个节点之后的那个时间点执行。

2、在上面两种情况，代码能够获取到的信息是不一样的。


### 迭代遍历的核心要点


前序

根左右
# 看这个方法，需要先把 right 放进去，再把 left 放进去。 
```
        if(root!=null)
          	stack.push(root);
        while(stack.size()!=0) {
            TreeNode p = stack.pop();
            list.add(p.val);
            if(p.right!=null)
              	stack.push(p.right);
            if(p.left!=null)
            		stack.push(p.left);
        }
        return list;
    }
```

# 后序遍历， 先把 left 放进去，再把 right 放进去
```
	public void postorderTraversal1(TreeNode root) {
		if(root.left != null)
			postorderTraversal(root.left);
		if(root.right != null)
			postorderTraversal(root.right);
		list.add(root.val);
```


### 经典做法

func inorderTraversalWithStack(root *TreeNode)[]int{
	var result []int

	stack := []*TreeNode{}
	cur := root

	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 已经访问当前节点了
			result = append(result, cur.Val)
			// 为什么？下面两行是必须的？
			cur = cur.Right
			continue
		}
	stack = append(stack, cur)
	cur = cur.Left
	}
	return result
}

