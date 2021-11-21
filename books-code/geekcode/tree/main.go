package tree

import "fmt"

// 树
// 节点的高度 = 节点到叶子节点到最长路径
// 节点到深度 = 根节点到这个节点所经历到边个树
// 节点到层数 = 节点到深度 + 1
// 树的高度 = 根节点的高度

// 二叉树的遍历时间复杂度是多少？ O(N)
// 相关情况

// 1.是卡特兰数，【卡特兰数】
// 是C[n,2n] / (n+1)种形状，c是组合数，节点的不同又是一个全排列，
// 一共就是n!*C[n,2n] / (n+1)个二叉树。
// 可以通过数学归纳法推导得出。
// 2.层次遍历需要借助队列这样一个辅助数据结构。(其实也可以不用，这样就要自己手动去处理节点的关 系，代码不太好理解，好处就是空间复杂度是o(1)。
// 不过用队列比较好理解，缺点就是空间复杂度是o(n) )。根节点先入队列，然后队列不空，取出对头元素，
// 如果左孩子存在就入列队，否则什么也不做，右孩 子同理。直到队列为空，则表示树层次遍历结束。
// 树的层次遍历，其实也是一个广度优先 BFS 的遍历算法。 [ 76赞]

// 二叉查找树
// 插入: 从根节点开始，依次比较要插入的数据和节点之间的大小关系
// 删除: 1、删除的节点没有子节点，直接删除
//      2、删除的节点只有一个子节点，更新父节点中，指向要删除节点的指针
//      3、如果要删除的节点有两个子节点，找到最小的最小节点，把它替换到要删除的节点上
// 需要两个字段记录📝：p, pp 记录当前要删除的节点，和父节点。
// 二叉查找树。它支持快速地查找、插入、删除操作。
//

//public void delete(int data){
//	Node p = tree;
//	Node pp = null;
//	while(p != null && p.data != data){
//		pp = p;
//		if (data > p.data) p = p.right;
//		else p = p.left
//	}
//
//	// 没有找到
//	if (p == null) return;
//	// 两个节点都存在的情况
//	if (p.left != null && p.right != null){
//		Node minP = p.right;
//		Node minPP = p; // minPP 表示 minP 的父节点
//		while (minP.left != null){
//			minPP = minP;
//			minP = minP.left
//		}
//	}
//	p.data = minP.data;
//	p = minP;
//	pp = minP;
//}

// 需要翻译成 GO 语言的代码。
// 叶子节点删除算法
// public void delete(int data){
//	Node p = tree;
//	Node pp = null;
//	while (p != null && p.Data != data){
//		pp = p;
//		if (data > p.data) p = p.right;
//		else p = p.left;
//	}
//	if (p==null) return;
//
//	if ( p.left !+ null && p.right != null) {
//		Node minP = p.right;
//		Node minPP = p;
//		while (minP.left != null){
//			minPP = minP;
//			minP = minP.left;
//		}
//		p.data = minP.data;
//		p = minP;
//		pp = minPP;
//	}
//	Node child;
//	if (p.left != null) child = p.left;
//	else if (p.right != null) child = p.right;
//	else child = null;
//	if (pp == null) tree = child;
//	else if (pp.left == p) pp.left = child;
//	else pp.right = child;
//}

// AVL 树🌲
// 平衡二叉树：二叉树中任意一个节点的左右子树的高度相差不能大于1。
// 这个条件对于工程实践来说，太严格了
// 为什么要求平衡？因为树高相差太多，二叉树就会退化为链表，各种操作的效率都急剧下降
//
// 平衡二叉查找树🌲：Splay Tree （伸展树🌲）、Treap 树堆
// 红黑树
// 1、根节点都是黑色的
// 2、每个叶子节点都是黑色的空节点（NIL）、叶子节点不存储数据
// 3、任何相邻的节点都不能同时为红色，也就是红色节点是被黑色节点隔开的
// 4、每个节点，从该节点达到其可达节点的所有路径，都包含相同数据的黑节点。
// 由来、特性、使用场景、已经能解决什么问题
// https://blog.csdn.net/abcdef314159/article/details/77193888
//

//
//栈、队列、散列表、跳表、树都是抽象的数据结构，它们在内存中存在的形式都要依赖于数组或者链表。
//如果用链表实现，很明显它们是动态的数据结构;
// 如果用数组实现，那么在扩容的时候会创建更大内存的 数组，原数组被废弃，从抽象角度看，它们仍旧是动态的
// 动态态数据结 构是支持动态的更新操作，里面存储的数据是时刻在变化的，通俗一点讲，
// 它不仅仅支持查询，还支持删 除、插入数据。而且，这些操作都非常高效。
//
//

// 借助树🌲来求解递归算法的时间复杂度
// 递归的思想就是，将大问题分解为小问题来求解，然后再将小问题分解为小小问题。这样一
// 层一层地分解，直到问题的数据规模被分解得足够小，不用继续递归分解为止
//

// 计算时间复杂度
// 只知道这颗树的高度 + 每层的消耗的时间

// 分析全排列的时间复杂度
//

// 全排列问题
// 利用递归树，分析全排列问题。
//
func printPerMutations(data []int, n int, k int) {
	if k == 1 {
		for i := 1; i < n; i++ {
			fmt.Println(data[i], " ")
		}
	}
	for i := 1; i < k; i++ {
		tmp := data[i]
		data[i] = data[k-1]
		data[k-1] = tmp

		printPerMutations(data, n, k-1)

		// swap back, 换回来
		tmp = data[i]
		data[i] = data[k-1]
		data[k-1] = tmp
	}
	//	suffixarray.New()
}

// 1 个细胞的生命周期是 3 小时，1 小时分裂一次。
// 求 N 小时后，容器内有多少细胞？
// 递归树🌲的方法进行分析。

// 12-30
// 这一周，专注于树算法的方方面面的实现与应用
// https://www.youtube.com/watch?v=IpyCqRmaKW4&list=PLqM7alHXFySHCXD7r1J0ky9Zg_GBB1dbk
// 树算法系列2

// 搜索，是树的最重要目的。
// 而搜索的效率，是由于树的高度来决定的。
// 因此，如果能够控制树的高度， 那么我们就可以提高树的搜索效率。
// AVL 树正是如此。

// 目标1
// 三种遍历的非递归写法，前中后，加上一个【求树深度的算法】
// https://seanlee97.github.io/2018/05/23/%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%9B%9B%E7%A7%8D%E9%81%8D%E5%8E%86%E6%96%B9%E5%BC%8F%E9%80%92%E5%BD%92%E9%9D%9E%E9%80%92%E5%BD%92%E5%AE%9E%E7%8E%B0/
// 遍历线索树🌲
//

// 求树高的非递归算法
// https://www.geeksforgeeks.org/iterative-method-to-find-height-of-binary-tree/
// 层次遍历的反向应用
//

// TODO: level travel Traversal
// 1、Create an empty queue q and an empty stack s
// 2、Enqueu the root node in the queue
// 3、Loop while queue is not NULL
// a) Dequeue while queue is not NULL
// b）Push the root node to the stack
// c) Enqueue root node's children (first right then left children to q
// d) Now pop all items from stack one by one and print them

// TODO: count half nodes in Binary Tree | GeeksforGeeks
//
