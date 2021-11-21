package tree


// https://www.youtube.com/watch?v=nPtARJ2cYrg
// 这个视频讲的非常清晰
// 思考问题的方法
// 1、我们有没有做过类似的题目
// 2、题目归类
// 3、这些类型的题目有什么解题技巧
// 4、能不能转化这个题目到相关的题？
// 5、这个难点在哪里。

// 明确事实（知识点）：tree is acyclic connected graph
// 我们将会使用 BFS
// 转化的难点是：我们不能 access 到这个节点的[父亲]
// 用什么技巧解决这个难点？  hashtable
// 如此一来题目就转化为：给定一个图和图中的一个节点，输出与此节点距离为 2 的所有节点。（BFS)



// 一口气写的
// 但是没有通过编译...
// 能够通过，但是不知道在哪里出错了
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if root == nil {
		return nil
	}

	m := map[*TreeNode]*TreeNode{}
	var dfs func(root, parent *TreeNode)
	dfs = func(root, parent *TreeNode){
		if root == nil {
			return
		}
		m[root] = parent
		dfs(root.Left, root)
		dfs(root.Right, root)
	}
	dfs(root, nil)

	var ret []int
	var search func(root *TreeNode)


	seen := map[*TreeNode]struct{}{}
	search = func(root *TreeNode){
		if root == nil {
			return
		}
		depth:=0
		queue := []*TreeNode{root}

		for len(queue) != 0 {
			depth++
			size := len(queue)
			for i:=0;i<size; i++{
				top := queue[0]
				queue = queue[1:]
				seen[top] = struct{}{}
				if K == depth {
					ret = append(ret, top.Val)
				}
				if top.Left != nil {
					queue = append(queue, top.Left)
				}
				if top.Right != nil {
					queue = append(queue, top.Right)
				}
				mm, ok := m[top]
				_,s  := seen[mm]
				if ok && mm != nil && !s{
					queue = append(queue, mm)
				}
			}
		}

	}
	search(target)
	return ret
}


// 参考答案的
// 学一下 map 里面如何判断元素存不存在的方法
// visited 这个接口体
func distanceKV2(root *TreeNode, target *TreeNode, K int) []int {
	data := make(map[*TreeNode]*TreeNode)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			data[node.Left] = node
		}
		if node.Right != nil {
			data[node.Right] = node
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	var ans []int
	visited := make(map[*TreeNode]bool)
	var queue []*TreeNode
	queue = append(queue, target)
	var depth int
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]
			if depth == K {
				ans = append(ans, node.Val)
				continue
			}
			parent := data[node]
			if parent != nil && !visited[parent] {
				queue = append(queue, parent)
			}
			if node.Left != nil && !visited[node.Left] {
				queue = append(queue, node.Left)
			}
			if node.Right != nil && !visited[node.Right] {
				queue = append(queue, node.Right)
			}
			visited[node] = true
		}
		depth++
	}
	return ans
}