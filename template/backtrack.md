### 回溯、backtracking、动态规划


### 算法核心代码

```
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return
    
    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
```
= 其核心就是 for 循环里面的递归，在递归调用之前「做选择」，在递归调用之后 【撤销选择】

```
for 选择 in 选择列表:
    # 做选择
    将该选择从选择列表移除
    路径.add(选择)
    backtrack(路径, 选择列表)
    # 撤销选择
    路径.remove(选择)
    将该选择再加入选择列表
```
三个重点：

- 选择列表
- 路径
- 撤销路径

### 全排列问题

变种：
1、不包含重复的数字。（简单）
2、包含重复数组。

### 多叉树

```
void traverse(TreeNode root) {
    for (TreeNode child : root.childern)
        // 前序遍历需要的操作
        traverse(child);
        // 后序遍历需要的操作
}
```








