package tree


// 正确率逐渐上来了...
// 终于找到做 算法题的相关技巧和实践功能了
//
func kthSmallest(root *TreeNode, k int) int {
	// 中序遍历
	// 第 N 个中序遍历的节点就是了//

	if root == nil || k == 0 {
		return 0
	}

	count := 0
	cur := root
	stack := []*TreeNode{}

	for len(stack) != 0 || cur != nil {
		if cur == nil {
			count ++
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if count == k {
				return cur.Val
			}

			cur = cur.Right
			continue
		}

		stack = append(stack, cur)
		cur = cur.Left
	}

	// k > count
	return 0
}


// 递归版的能不能实现一下？中序遍历的递归模式
// 用闭包函数，引用两个全局变量
// ...
//


// 都 ac 了, 但是效率还是不够高...
func kthSmallestRecursion(root *TreeNode, k int) int{
	if root == nil || k == 0{
		return 0
	}

	count := 0
	var ret int
	var helper func(root *TreeNode)
	helper = func(root *TreeNode){
		if root == nil {
			return
		}
		helper(root.Left)
		count++
		if count == k {
			ret = root.Val
		}
		helper(root.Right)
	}

	helper(root)
	return ret
}


// 迭代遍历思想
// https://leetcode.com/problems/kth-smallest-element-in-a-bst/discuss/63723/Golang-iterative
// 差不多，但是代码没有我的清晰
//



// 递归思想
// 当时代码有点丑...
// 太多引用符号了吧，
// 有可能会导致出问题，
// 因为在多核 CPU 如果不上锁的话
//

