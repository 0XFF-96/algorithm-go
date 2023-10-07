package tree



// 看一下相关的视频，看看有没有相关启发...
// https://www.youtube.com/watch?v=PQKkr036wRc
//


// 什么是 TreeMap ?
// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/discuss/231148/Java-TreeMap-Solution


// 遍历的顺序没对
// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/discuss/?currentPage=1&orderBy=most_relevant&query=golang

import "sort"
// 没有AC
// 有相关的题目...
// 顺序的问题没有搞定过...
// [0,8,1,null,null,3,2,null,4,5,null,null,7,6]
// [0,8,1,null,null,3,2,null,4,5,null,null,7,6]

// [[8],[0,3,6],[1,4,5],[7,2]]
// [[8],[0,3,6],[1,4,5],[2,7]]
// 有个遍历顺序不对
// 比较诡异...

// 如果在同一个垂直距离上，应该怎么输出？
func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	verticalMap := map[int][]int{}

	var helper func(root *TreeNode, ver int)

	helper = func(root *TreeNode, ver int){
		if root == nil {
			return
		}
		_, ok := verticalMap[ver]
		if !ok {
			verticalMap[ver] = []int{root.Val}
		} else {
			verticalMap[ver] = append(verticalMap[ver], root.Val)
		}
		helper(root.Left, ver-1)
		helper(root.Right, ver+1)
	}

	helper(root, 0)

	// sort map by key order ...
	//
	// tree map
	var keys []int
	for k, _ := range verticalMap{
		keys = append(keys, k)
	}
	sort.Ints(keys)


	var ret [][]int
	for _, k := range keys {
		// deep copy
		ret = append(ret, verticalMap[k])
	}
	return ret
}


// 正确答案是这个
// 没有搞明白怎么回事
// ...
func verticalTraversalV2(root *TreeNode) [][]int {
	// store in maps
	m := make(map[int]map[int][]int)
	var dfs func(root *TreeNode, x int, y int)
	dfs = func(root *TreeNode, x int, y int) {
		if root == nil { return }
		if m[x] == nil {
			m[x] = make(map[int][]int)
		}
		if m[x][y] == nil {
			m[x][y] = make([]int, 0)
		}
		m[x][y] = append(m[x][y], root.Val)
		dfs(root.Left, x-1, y+1)
		dfs(root.Right, x+1, y+1)
	}
	dfs(root, 0, 0)

	// sort x keys to ensure the order
	keysX := make([]int, len(m))
	idx := 0
	for k, _ := range m {
		keysX[idx] = k
		idx++
	}
	sort.Ints(keysX)
	res := [][]int{}
	for i:=0; i<len(keysX); i++ {
		// sort y keys to ensure the order
		keysY := make([]int, len(m[keysX[i]]))
		idx := 0
		for k, _ := range m[keysX[i]] {
			keysY[idx] = k
			idx++
		}
		sort.Ints(keysY)

		y := []int{}
		for j:=0; j<len(keysY); j++ {
			sort.Ints(m[keysX[i]][keysY[j]])
			y = append(y, m[keysX[i]][keysY[j]]...)
		}
		res = append(res, y)
	}
	return res
}


// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
//


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 抄答案 AC 了
// 但不知道怎么回事.
func verticalTraversalV3(root *TreeNode) [][]int {
	// x:y value
	// m := ma[int]map[int][]int{}{}

	// x, y value ?
	m := make(map[int]map[int][]int)
	var dfs func(root *TreeNode, x int, y int)

	dfs = func(root *TreeNode, x int, y int){
		if root == nil {
			return
		}
		if m[x] == nil {
			m[x] = make(map[int][]int)
		}
		if m[x][y] == nil {
			m[x][y] = make([]int,0 )
		}
		m[x][y] = append(m[x][y], root.Val)

		dfs(root.Left, x-1, y+1)
		dfs(root.Right,x+1, y+1)
	}


	dfs(root, 0, 0)

	// sort x keys to ensure the order
	keysX := make([]int, len(m))
	idx := 0
	for k, _ := range m {
		keysX[idx] = k
		idx++
	}
	sort.Ints(keysX)

	var ret [][]int

	for i:=0; i < len(keysX);i++{

		keysY := make([]int, len(m[keysX[i]]))
		idx := 0
		for k,_ := range m[keysX[i]] {
			keysY[idx] = k
			idx++
		}
		sort.Ints(keysY)

		y := []int{}
		for j:=0; j < len(keysY);j++{

			// 这次是为何？
			sort.Ints(m[keysX[i]][keysY[j]])
			y = append(y, m[keysX[i]][keysY[j]]...)
		}

		ret = append(ret, y)
	}

	return ret
}