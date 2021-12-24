一、面试题50 树中两个结点的最低公共祖先 （*) ：
 soloutiion: 先求出两个结点到根结点的路径，然后从路径中找出最后一个公共结点

二、把字符串转化为整数：
  ord(), 通过 ASCIID 码进行转换
  ```
      for k, s in enumerate(string):
        if s.isdigit():  # 数字直接运算
            val = ord(s) - ord('0')
            ret = ret * 10 + val
  ```

三、实现单例的方式 py， golang , sync.once 调用一次..:
  ```
  class SingleTon(object):
    _instance = {}

    def __new__(cls, *args, **kwargs):
        if cls not in cls._instance:
            cls._instance[cls] = super(SingleTon, cls).__new__(cls, *args, **kwargs)
        # print cls._instance
        return cls._instance[cls]
  ``` 有元信息

四、golang 单例创建的模式：https://blog.csdn.net/qibin0506/article/details/50733314 
  1、能否做到线程安全：https://blog.csdn.net/lengyuezuixue/article/details/78893548 
  2、比较好的一种方式sync.Once


五、题目：
二维数组中，每行从左到右递增，每列从上到下递增，给出一个数，判断它是否在数组中
思路：从左下角或者右上角开始比较

六、替换空格 : 直接使用Python字符串的内置函数, 使用正则表达式 

七、从尾到头打印单链表 ： 一种使用栈， 二、直接递归
   递归的方法：
        def print_link_recursion(links):
          if links:
              print_link_recursion(links.next)
              print links.val 

八、重建二叉树: 
  思路：前序的第一个元素是根结点的值，在中序中找到该值，中序中该值的左边的元素是根结点的左子树，右边是右子树，然后递归的处理左边和右边。 

  ```python3
  def construct_tree(preorder=None, inorder=None):
    """
    构建二叉树
    """
    if not preorder or not inorder:
        return None
    index = inorder.index(preorder[0])
    left = inorder[0:index]
    right = inorder[index+1:]
    root = TreeNode(preorder[0])
    root.left = construct_tree(preorder[1:1+len(left)], left)
    root.right = construct_tree(preorder[-len(right):], right)
    return root
  ```

九、用两个栈实现队列 
  思路：要求：用两个栈实现队列，分别实现入队和出队操作 思路：一个栈负责入队，另一个负责出队，出栈为空则从入栈中导入到出栈中。

  ```
  class MyQueue(object):
    def __init__(self):
        self.stack = []
        self.stack2 = []

    def push(self, val):
        self.stack.append(val)

    def pop(self):
        if self.stack2:
            return self.stack2.pop()
        while self.stack:
            self.stack2.append(self.stack.pop())
        return self.stack2.pop() if self.stack2 else u'队列为空'
  ```


十、旋转数组的最小数字：使用二分法，但要考虑[1, 0, 0, 1]这种数据，只能顺序查找。
```
  def find_min(nums):
    if not nums:
      return false
    length = len(nums)
    left, right = 0, length -1
    while nums[left] >= nums[right]:
      if right - left == 1:
        return nums[right]
      mid = (left + right) / 2

      if nums[left] == nums[mid] == nums[right]:
        return min(nums)
      
      if nums[left] <= nums[mid]:
        left = mid 
      if nums[right] >= nums[mid]:
        right = mid 
    return nums[0]
```


十一、斐波那契数列
  def fib(num):
    a, b =0, 1
    for i in range(num):
      yield b 
      a, b = b, a + b

十二、二进制中1的个数:   n = n & n-1 公式， ret = 0 计数器...

十三、(不会）要求：求一个数的整数次方
    思路：需要考虑次方是正数、负数和0，基数是0
    1、https://github.com/JushuangQiao/Python-Offer/tree/master/third/third

十四、O(1)时间删除链表结点

```
  def delete_node(link, node):
    if node == link:
      del node
    if node.next is None: # 尾节点...
      while link:
        if line.next == node:
          link.next = None
        link = link.next 
    else:
      node.val = node.next.val
      n_node = node.next 
      node.next = n_node.next 
      del n_node

十五、调整数组顺序使奇数位于偶数前面 
思路：使用两个指针，前后各一个，为了更好的扩展性，可以把判断奇偶部分抽取出来

def reorder(nums, func): 乱序数组， 判断偶的函数, 将判断奇偶数的部分抽出来..
```

十六、链表中的倒数第 k 个结点:
  快慢指针
```
def last_kth(link, k):
  if not link or k <= 0:
    return None
  move = link 
  while move and k-1 >= 0:
    move = move.next 
    k -= 1 
  while move:
    move = move.next 
    link = link.next 
  if k == 0:
    return link.val
  return None

十七、翻转链表: 需要考虑空结点和只有一个节点的情况

def reverse_link(head):
  if not head or not head.next:
    return head
  
  then = head.next 
  head.next = None 
  last = then.next 

  while then:
    then.next = head 
    head = then 
    then = last 
    if then:
      last = then.next 
  return head 


十八、合并两个排序的链表：
```
def merge_link(head1, head2):
    if not head1:
        return head2
    if not head2:
        return head1
    if head1.val <= head2.val:
        ret = head1
        ret.next = merge_link(head1.next, head2)
    else:
        ret = head2
        ret.next = merge_link(head1, head2.next)
    return ret
```

判断树的子结构：
思路： 使用递归： sub_tree(tree1.left, tree2.left) and sub_tree(tree1.right, tree2.right)
```
def sub_tree(tree1, tree2):
    if tree1 and tree2:
        if tree1.val == tree2.val:
            return sub_tree(tree1.left, tree2.left) and sub_tree(tree1.right, tree2.right)
        else:
            return sub_tree(tree1.left, tree2) or sub_tree(tree1.right, tree2)
    if not tree1 and tree2:
        return False
    return True
```

二十、面试题21 包含min函数的栈：
要求：栈的push，pop，min操作的时间复杂度都是O(1)
思路：使用一个辅助栈保存最小值

二十一、栈的压入弹出序列： 模拟栈操作...
要求：判断给定的两个序列中，后者是不是前者的弹出序列，给定栈不包含相同值
思路：使用一个辅助栈, 如果辅助栈栈顶元素不等于出栈元素，则从入栈中找改值，直到入栈为空
如果最后出栈序列为空，则是入栈的弹出序列值

https://www.cnblogs.com/edisonchou/p/4779755.html 思路..


二十三：从上往下打印二叉树:
层次遍历
```
  def bfs(tree):
    if not tree:
      return None
    stack = [tree]
    ret = []
    while stack:
      node = stack.pop(0)
      ret.append(node.val)
      if node.left:
        stack.append(node.left)
      if node.right:
        stack.append(node.right)
    return ret
```

### 二十四：二叉树中和为某一值的路径 
  1、要求：输入一棵二叉树和一个值，求从根结点到叶结点的和等于该值的路径
  2、深度优先搜索变形

  数据结构： path, sums, ret 结构 数值..是否相等...
  一深度遍历
    左子树存在，递归左子树，加入数值..
    右子树存在，递归右子树，加入数值..