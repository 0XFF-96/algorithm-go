### 回溯、backtracking、动态规划

review:1

### 大纲

### 回溯算法能够解决什么问题？

组合问题：N 个数里面按一定规则找出k个数的集合
切割问题：一个字符串按一定规则有几种切割方式
子集问题：一个 N 个数的集合里有多少符合条件的子集
排列问题：N 个数按一定规则全排列，有几种排列方式
棋盘问题：N 皇后，解数独等等

### 需要思考的问题
解决一个回溯问题，实际上就是一个多叉树的遍历过程

- 路径：也就是已经做出的选择。
- 选择列表：也就是你当前可以做的选择。
- 结束条件：也就是到达决策树底层，无法再做选择的条件。
或者满足了题目的基本要求的条件。

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


### 算法核心代码（自然语言描述，伪代码描述）

void backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}


### 如何思考多叉树的时间和空间复杂度？

- 树高的深入
- 

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


### 核心要点

func subsets(nums []int) [][]int {
    // 保存最终结果
    result := make([][]int, 0)
    // 保存中间结果
    list := make([]int, 0)
    backtrack(nums, 0, list, &result)
    return result
}

// nums 给定的集合
// pos 下次添加到集合中的元素位置索引
// list 临时结果集合(每次需要复制保存)
// result 最终结果
func backtrack(nums []int, pos int, list []int, result *[][]int) {
    // 把临时结果复制出来保存到最终结果
    ans := make([]int, len(list))
    copy(ans, list)
    *result = append(*result, ans)
    // 选择、处理结果、再撤销选择
    for i := pos; i < len(nums); i++ {
        list = append(list, nums[i])
        backtrack(nums, i+1, list, result)
        list = list[0 : len(list)-1]
    }
}





