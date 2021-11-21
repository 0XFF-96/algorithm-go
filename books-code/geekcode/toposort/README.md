// 拓扑排序, graph
// Topologicial Sort
// build system
//
// 1、可以用深度搜索的思想实现拓扑排序：https://blog.csdn.net/lisonglisonglisong/article/details/36896191
// 2、
// 3、尝试练习
// 4：Topological Sort, Idea for a Directed Acyclic (DAG) is a linear ordering
// of vertices such that for every directed edge(u,v), vertex u comes before v in the
// ordering
// 应用
// 1、Build Systems, 2、Advanced-Packaging Tool 3、Task Scheduling 4、Pre-requisite Problems
//

// void Graph::TopologicalSortUitl(int v, bool visited[], stack<int> &stack){
//
//	// Mark the current node as visited.
//	// Recur for all the vertices adjacent to this vertex
//
//	list<int>::iterator i;
//	for (i = adj[v].begin(); i != adj[v].end(); ++i){
//		if (!visisted[*i]){
//			topolocialSortUtil(*i, visited, Stack)
//		}
//
//	// Push current vertext to stack which stores result
//	Stack.push(v);
//	}
//}
//
// void Graph::topolocalSort(){
//
//	stack<int> Stack;
//
//	Mark all the vertices as not visited
//	bool *visited = new bool[V];
//	for (int i = 0; i < V; i++){}
//		visited[i] = false
//	}
//	// Call the recursive helper function to store Topological
//	// Sort starting from all vertices one by one
//	for (int i =0;i < V;i++){
//		if (visited[i] == false){
//			topologicalSortUtil(i, visted, stack);
//		}
//
//	// 输出
//	// Print contents of stack
//	while (Stack.empty() == false){
//		cout << Stack.top() << " ";
//		stack.pop();
//	}
// }
//

// 前置知识拓展相关...
// https://www.youtube.com/watch?v=-mOEd_3gTK0
// 最短路径算法
