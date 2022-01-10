
### 算法题目回顾

1、买股票的最佳时机【❌】一、（最大子序列之和）二、动态规划解法乘积最大子序列问题。 最长公共子序列。【最长子串，这样就需要用指针解放！！！】
1、做一个二维数组，然后学会初始化
2、当相同时，斜上角的 + 1。 当不同时，两个元素，对比，进行最大值。
2、jump game , https://leetcode-cn.com/problems/jump-game/ 。 3、跳跃游戏，看最远能到的地方。 rightmost 的问题。 https://leetcode-cn.com/problems/jump-game/
3、微信发红包， 假设一共有 N 元，一共有 K 个人，则可以每个人拿到的钱为 random(N - (K - 1) * 0.01)，然后更新N，直到最后一个人就直接 N。
4、合并有序数组， 利用双指针法。我们知道如果从头开始，那么数组 A 的元素会被覆盖，那么我们可以从后面开始 。 
链表合并问题， k 个链表合并问题。 https://leetcode-cn.com/problems/merge-k-sorted-lists/solution/4-chong-fang-fa-xiang-jie-bi-xu-miao-dong-by-sweet/ 。 5、.剪绳子【贪心】【❌】【低优先级】
6、7.找出不重复的元素个数。一个数据先递增再递减，找出数组不重复的个数。不能使用额外空间，复杂度o(n)
8、【找范围】【桶排序】高考成绩2000万数据，分数0-750，如何快速知道你的排名，如何知道任一分数排名？
9、链表相交【❌】【❌】【重点】一定会到达这样一个节点，在这个节点两个指针走过的路长分别为 AC + CD + BC 和 BC + CD + AC，10、互相关注表设计。 场景题：需求：谁关注了我，我关注了谁，谁与我互相关注。表该如何设计，索引怎么建。查询语句怎么写
个人见解，表可以有：id、粉丝id、被追随者id 三个字段，索引可以建两个，分别是粉丝 id 和被追随者 id。
观点1: 需要建立两个表。 观点2: 是否需要缓存？11、[103. 二叉树的锯齿形层次遍历](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/) 。 二叉树遍历方法三种。 （递归和迭代）（右视图）层序遍历的时候，判断是否遍历到单层的最后面的元素，如果是，就放进result数组中，随后返回result就可以了。12、链表求和。 遍历，并用一个标志位来表示进位即可。
注意⚠️，数字的表示是，顺序还是逆序存放的呢。
进阶：假设这些数位是正向存放的，请再做一遍。
（先反转数组， 再求和一遍） 13、生成随机数。 【❌】【❌】【❌】给定一个 0-4 随机数生成器 如何生成 0-6 随机数。14、19.二叉树的最近公共祖先【❌】lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q)。 15、最大路径和。 【❌】【❌】路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，【且不一定经过根节点】。（ 二叉树的路径和）。【且不一定经过跟节点】maxGain(), 需要用必包函数。 16、二叉树后序遍历【❌】【❌】https://leetcode.com/problems/binary-tree-postorder-traversal/discuss/331618/Golang-solution-(Iterations) 。 这里要增加一个辅助节点来节点上一次放入结果数组的值，当一个节点左右都是空的时候，就可以放入结果集，当上一个放入结果集的节点是他的孩子节点的时候，说明他的孩子已经访问完成了，到这里就是第三次了就可以放入了，还有一点要注意的是，当一个节点的左右不为空时，要先加入右孩子，再加入左孩子，这样才能先访问左孩子。
17、接雨水【❌】
18、反转矩阵和螺旋矩阵。 【❌】
19、重建二叉树。题目很像 从中序和后续，得到后序遍历的序列。 【❌】20、二叉树的路径总和。 【✅】
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
https://leetcode-cn.com/problems/path-sum-ii/ 。21、区间合并问题【❌】【❌】https://leetcode-cn.com/problems/merge-intervals/21、归并排序和二分查找排序。 二分查找就是将查找有序的范围不断均分，将查找的值和中间值进行比较，从而缩小范围。
如果在链表中使用二分查找，那就得用跳表了。另外构建一个链表，链表有两个指针，一个是 next，指向下一个节点，一个是 main，指向原链表。保持每两个主链表节点有一个跳表节点。如 redis 的 zset 。
 22、版本比较问题。版本数字比较。版本数字比较，比如"1.10.0"版本比"1.8.1"版本新，不允许使用 split 等函数。 双指针。 23、文件拷贝。【❌】24、缺失的第一个正数。给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
用位图法，遍历数组将值为 i 的数字转移到下标为 i-1 的位置。再次遍历查看不对的值即可。
25、有序数组中，第一个位置和最后一个位置。 https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/ 。 
26、二叉树右视图。27、三数之和。【排序 + 双指针】【有点利用 twoSum 的相关思想了】28、最长不含重复字符的子字符串。 https://github.com/lewiscrow/WorkHardAndFindJob/blob/master/%E5%A4%8D%E4%B9%A0/%E9%9D%A2%E8%AF%95/%E6%89%8B%E6%92%95%E5%AD%97%E8%8A%82%E8%B7%B3%E5%8A%A8%E9%9D%A2%E8%AF%95%E6%97%B6%E5%87%BA%E7%8E%B0%E8%BF%87%E7%9A%84%E7%AE%97%E6%B3%95%E9%A2%98.md 。29、求中位数【❌】30、和为K的子数组 。 【前序和技巧】给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
https://leetcode-cn.com/problems/subarray-sum-equals-k/ 。31、复制复杂链表【✅】 
方法一：哈希表。方法二：原地合并，然后再分开链表。 32、两个有序数组的中位数。 简单粗暴，先将两个数组合并，两个有序数组的合并也是归并排序中的一部分。然后根据奇数，还是偶数，返回中位数。不需要合并两个有序数组，只要找到中位数的位置即可。由于两个数组的长度已知，因此中位数对应的两个数组的下标之和也是已知的。维护两个指针，初始时分别指向两个数组的下标 00 的位置，每次将指向较小值的指针后移一位（如果一个指针已经到达数组末尾，则只需要移动另一个数组的指针），直到到达中位数的位置。 33、rand7() 实现 rand10() 【❌】（生成 0 和 1） 【左程云老师】(rand_X() - 1) × Y + rand_Y() (rand(X)-1)*Y + randY() -----生成 [1,X*Y]区间，并且等 。
此题范围1-7，可以划分为生成1,2,3时返回0，生成5,6,7返回1，随机到4就重新生成。 1-10的范围可以表示成1 + (0到9) 二进制的0000-1111范围为0-15，如果范围超过了9就重新生成。 一定可以返回一个0-9的数。func rand10() int {
    for {
        row := rand7()
        col := rand7()
        idx := (row-1)*7 + col
        if idx <= 40 {
            return 1 + (idx-1)%10
        }
    }
}https://leetcode-cn.com/problems/implement-rand10-using-rand7/solution/gong-shui-san-xie-k-jin-zhi-zhu-wei-shen-zmd4/ 。34、前 K 大数。 https://leetcode-cn.com/problems/kth-largest-element-in-an-array/solution/partitionfen-er-zhi-zhi-you-xian-dui-lie-java-dai-/ 【❌】【视频重点】36、求x的n次方。
37、两两交的链表。 28、旋转矩阵【❌】【操作 matrix 】 最直白蛮力的方法，先转置后镜像对称。 上下反转，再对角反转。 90 度的含义是这个。 找到关键等式子， matrix[col][n - row -1] = matrix[row][col] 。 row = col 
col = n - row - 1 临时遍历。 



### 不怎么熟悉的题目


需要用自然语言描述这些问题、需要用图形画出这些问题的核心解法、1、滑动窗口【❌】和双指针题目类似，更像双指针的升级版，滑动窗口核心点是维护一个窗口集，根据窗口集来进行处理
核心步骤，
，right 右移，什么时候停止， 收缩？，什么时候停止收缩，left 向右移动， 求结果。 按照题目的意思，将核心步骤迁移和细化为代码。
2、无重复字符的最长子串。3、排序算法（冒泡、归并、快排、插入、选择、堆排）拓扑排序（自然语言描述）？十大经典排序， http://cnblogs.com/onepixel/p/7674659.html 。 二叉堆的作用和意义。各种排序算法， https://leetcode-cn.com/problems/squares-of-a-sorted-array/solution/ge-chong-pai-xu-shuang-zhi-zhen-by-toxic-3/ 。 4、循环排序， Cyclic Sort5、多路归并（K-way merge）
6、在数组中找到第 k 大的元素 。 （快排（每次排序能够确定一个元素的最终位置、堆排序）7、字节跳动高频题之排序奇升偶降链表
， 奇偶链表排序。 https://zhuanlan.zhihu.com/p/311113031时间复杂度和常数级空间复杂度下，对链表进行排序。 找到中点断开，翻转后面部分，然后合并前后两个链表。给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。 要求返回这个链表的 深拷贝。8、水壶问题， https://zhuanlan.zhihu.com/p/79938638 。【❌】 9、给定一个二叉树，返回其节点值的锯齿形层次遍历。Z 字形遍历 。10、岛屿问题。 【❌】
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
11、单调队列结构。【好像多数人不知道quick-select。其实就是只利用partition进行部分排序。由于不需要全序数组，所以巧妙地将复杂度降到O(n)。但也诚如许多回答所说，如果输入更像是一个无尽的流（很大）的话，而所需要保持的集合（k）很小的话，更适合用堆。不管怎么说，quick-select都是很优秀的算法，毕竟不是什么时候都需要做大数据处理，那是分布式计算的领域。】
12、接雨水问题。 https://leetcode-cn.com/problems/trapping-rain-water/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-8/ 。 13、区间合并的问题 【❌】【❌】
需要看视频，补补。 14、岛屿数量 【❌】【没有思路】【并查集】【❌】上面大部分来自于：
https://zhuanlan.zhihu.com/p/381112418  。 15、异位字母词。 【❌】【有思路】https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/solution/438-zhao-dao-zi-fu-chuan-zhong-suo-you-z-nx6b/。 16、三数之和【✅】犹豫不决先排序，步步逼近双指针。two-sum 的问题。 

17、环状数组，是否存在循环。 【快慢指针】【图的遍历标记法】【什么是拓扑排序】【判断图中是否有环的问题】
问题转化。 【❌】
https://leetcode-cn.com/problems/circular-array-loop/18、任务调度器，区间问题。 【❌】https://leetcode-cn.com/problems/task-scheduler/solution/ren-wu-diao-du-qi-by-leetcode-solution-ur9w/ 。note1、从大数据的面试题目中，可以看到如何应对海量数据的处理。
2、https://hackernoon.com/the-2018-devops-roadmap-31588d8670cb  2028-devops-roadmap, 作为一个开发者应该学习些什么？3、50 道高频题目， http://bilibili.com/read/cv9852205 。 4、38 道高频题目， https://zhuanlan.zhihu.com/p/137835226 。 

### 频率击破

1、http://nowcoder.com/discuss/577995 。 [github.com/afatcoder/LeetcodeTop/blob/master/](https://github.com/afatcoder/LeetcodeTop/blob/master/bytedance/backend.md)
2、https://codetop.cc/home3、[https://github.com/lewiscrow/WorkHardAndFindJob/blob/master/复习/面试/手撕字节跳动面试时出现过的算法题.md](https://github.com/lewiscrow/WorkHardAndFindJob/blob/master/%E5%A4%8D%E4%B9%A0/%E9%9D%A2%E8%AF%95/%E6%89%8B%E6%92%95%E5%AD%97%E8%8A%82%E8%B7%B3%E5%8A%A8%E9%9D%A2%E8%AF%95%E6%97%B6%E5%87%BA%E7%8E%B0%E8%BF%87%E7%9A%84%E7%AE%97%E6%B3%95%E9%A2%98.md) 。 【过一遍这个算法题目】【题目】4、https://github.com/0voice/campus_recruitmen_questions5、https://mp.weixin.qq.com/s/7rCsaGy8B2lwbZ4duEC7Nw6、http://nowcoder.com/discuss/428774?type=2 。 7、https://github.com/youngyangyang04/leetcode-master 【挺不错的资料】【值得多看几遍】8、[https://programmercarl.com/0059.螺旋矩阵II.html#其他语言版本](https://programmercarl.com/0059.%E8%9E%BA%E6%97%8B%E7%9F%A9%E9%98%B5II.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC) 【值得多看】【总结得不错】prepared link:1、https://github.com/lewiscrow/WorkHardAndFindJob/tree/master/%E5%A4%8D%E4%B9%A0/%E9%9D%A2%E8%AF%952、https://github.com/KeKe-Li/data-structures-questions/blob/master/src/chapter05/golang.01.md#%E6%A0%88%E7%9A%84%E5%86%85%E5%AD%98%E6%98%AF%E6%80%8E%E4%B9%88%E5%88%86%E9%85%8D%E7%9A%843、https://github.com/CyC2018/CS-Notes/blob/master/notes/%E6%95%B0%E6%8D%AE%E5%BA%93%E7%B3%BB%E7%BB%9F%E5%8E%9F%E7%90%86.md#%E6%95%B0%E6%8D%AE%E5%BA%93%E7%B3%BB%E7%BB%9F%E5%8E%9F%E7%90%86 。挑一个项目说，挑一个模块说，具体的难点在哪里，怎么解决这个难点。当时有么有发现什么 Bug 。业界的解决办法是什么？

### 子序列问题


1、最长公共子前缀， https://leetcode-cn.com/problems/longest-common-prefix/solution/zui-chang-gong-gong-qian-zhui-by-leetcode-solution/ 有横向扫描法和纵向扫描法。难点在于，如何控制返回。
2、矩阵中的最长递增路径【❌】3、最长重复子数组【❌】记住，子序列默认不连续，子数组默认连续
dp[i][j]代表以A[i-1]与B[j-1]结尾的公共字串的长度,公共字串必须以A[i-1]，B[j-1]结束，即当A[i-1] == B[j-1]时，dp[i][j] = dp[i-1][j-1] + 1; 当A[i-1] != B[j-1]时，以A[i-1]和B[j-1]结尾的公共字串长度为0,dp[i][j] = 0。输出最大的公共字串的长度即为最长重复字串。 打个表会更直观一点。https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/solution/zui-chang-zhong-fu-zi-shu-zu-by-leetcode-solution/ 4、最长公共子序列。 https://leetcode-cn.com/problems/longest-common-subsequence/solution/zui-chang-gong-gong-zi-xu-lie-by-leetcod-y7u0/ 。 相等的情况下， dp[i][j] = dp[i - 1][j - 1] + 1; 。 
不相等的情况下，  dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]); 。 （子序列 != 连续。 子数组 = 连续。 ）5、最长连续序列。哈希表， 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。 https://leetcode-cn.com/problems/longest-consecutive-sequence/solution/ha-xi-zui-qing-xi-yi-dong-de-jiang-jie-c-xpnr/ 。 6、最长递增子序列。 https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/dong-tai-gui-hua-er-fen-cha-zhao-tan-xin-suan-fa-p/ 。定义 dp[i] 为，考虑前 i 个元素，以第 i 个数字结尾的最长上升子序列的长度，注意 nums[i] 必须被选取。 
核心代码。    for (int i = 0; i < nums.length; i++) {
        for (int j = 0; j < i; j++) {
            if (nums[i] > nums[j]) 
                dp[i] = Math.max(dp[i], dp[j] + 1);
        }
    }

### 排序算法

1、排序算法分类， 分为比较排序和非比较排序。 2、快速排序， 快速排序最优的情况就是每一次取到的元素都刚好平分整个数组。 3、堆排序。
4、计数排序 ->. 桶排序。
5、归并排序（重点）资料1、https://zhuanlan.zhihu.com/p/320419705 。 2、归并排序， http://cnblogs.com/chengxiao/p/6194356.html 。归并排序是稳定排序，它也是一种十分高效的排序，能利用完全二叉树特性的排序一般性能都不会太差。 3、http://cnblogs.com/chengxiao/p/6194356.html 。

