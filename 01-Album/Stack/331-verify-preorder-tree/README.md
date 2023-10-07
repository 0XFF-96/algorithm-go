
<!-- We just need to remember how many empty slots we have during the process.

Initially we have one ( for the root ).

for each node we check if we still have empty slots to put it in.

a null node occupies one slot.
a non-null node occupies one slot before he creates two more. the net gain is one -->


### 题目知识点
1. 树的出度和入度。
2. 前缀树的有效性。
3.

### 解题方法
1. 累乘，利用出度入度的相关知识
2. 栈，（中序遍历的思想）
3. 递归，