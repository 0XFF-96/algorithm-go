// Counting Indegree and Outdegree, SIMPLE & CLEAR!
// 
// 
func isValidSerialization(preorder string) bool {
    // binary tree
    //  Find an algorithm without reconstructing the tree
    // 中序遍历的思想，转变
    // 利用出度和入度的思想
    // 
    indegree := -1 
    strs := strings.Split(preorder, ",")
    for _, s := range strs {
        indegree++ // all nodes have 1 indegree (root compensated)
        
        // total indegree should never exceeds 0 
        if indegree > 0 {
            return false
        }
        if s != "#" {
            indegree -= 2
        }
    }
    return indegree == 0 
}