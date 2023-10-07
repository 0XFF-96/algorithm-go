package min_stack

// https://www.nowcoder.com/practice/4c776177d2c04c2494f2555c9fcc1e49?tpId=13&tqId=11173&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking
// 定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的min函数（时间复杂度应为O（1））
// 要求 O(1) 所以要借助辅助栈
//
//因此, 用另一个栈来保存这些元素是再合适不过的了. 我们叫它最小元素栈.
//每次压栈操作时,
//如果压栈元素比当前最小元素更小,
//就把这个元素压入最小元素栈, 原本的最小元素就成了次小元素.
//同理, 弹栈时, 如果弹出的元素和最小元素栈的栈顶元素相等, 就把最小元素的栈顶弹出
// https://studygolang.com/articles/19698

type MinStack struct {
	data []interface{}
	minIndex []int
}

func Constructor() MinStack{
	var instance MinStack
	return instance
}

func (m *MinStack) Push(num int){
	m.data = append(m.data, num)
	if len(m.minIndex) ==0{
		m.minIndex = append(m.minIndex,0)
	} else {
		if m.data[m.minIndex[len(m.minIndex) -1]].(int) > num {
			m.minIndex = append(m.minIndex, len(m.data)-1)
		}
	}
}

func (m *MinStack) Pop() {
	length := len(m.data)
	currentMinIndex := length -1

	if currentMinIndex == m.minIndex[len(m.minIndex) -1]{
		m.minIndex = m.minIndex[:len(m.minIndex)-1]
	}
	m.data = m.data[:len(m.data)-1]
}

// 栈的压入弹出顺序
// 模拟栈操作
// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。[假设压入栈的所有数字均不相等]。
// 如果数字相等？
// 判断的是否是pop次序的思路一样，但是你的程序写的太棒了，我借用辅助空间，写的一大堆代码。。。惭愧！
//class Solution {
//public:
//bool IsPopOrder(vector<int> pushV,vector<int> popV) {
//if(pushV.size() == 0) return false;
//vector<int> stack;
//for(int i = 0,j = 0 ;i < pushV.size();){
//stack.push_back(pushV[i++]);
//while(j < popV.size() && stack.back() == popV[j]){
//stack.pop_back();
//j++;
//}
//}
//return stack.empty();
//}
//};
// https://blog.csdn.net/pbrlovejava/article/details/103133279
// 尝试写一下关于这个的测试用例
