package queueUsingStack

import (
	"testing"
)

// https://studygolang.com/articles/10481


type stack []int
type Queue struct {
	inStack []int
	outStack []int
}

func (q *Queue)Pop() (top int){
	oLen := len(q.outStack)
	iLen := len(q.inStack)
	//if oLen == 0 && iLen == 0 {
	//	return 0
	//}
	if len(q.outStack) != 0 {
		top = q.outStack[oLen -1]
		q.outStack = q.outStack[:oLen-2]
		return top
	} else if iLen != 0{
		q.outStack = append(q.outStack, q.inStack...)
		top = q.outStack[len(q.outStack)-1]
		q.outStack = q.outStack[:len(q.outStack)-2]
		return
	} else {
		return 0
	}
}

func (q *Queue) Push(n int){
	q.inStack = append(q.inStack, n)
}

// 栈A用来作入队列
// 栈B用来出队列，当栈B为空时，栈A全部出栈到栈B,栈B再出栈（即出队列）
func TestTwoStackToQueue(t *testing.T){

}

// 用两个栈实现队列
// ...
// https://studygolang.com/articles/10481
// https://hansedong.github.io/2019/04/02/15/
// golang 栈的基本使用
// 旋转数组最小数字
// 1、利用冒泡的思想，比较 n2
// 2、排序  nlogn
// 3、二分搜索 logn
// https://www.nowcoder.com/questionTerminal/9f3231a991af4f55b95579b44b7a01ba?answerType=1&f=discussion
// 注意此数组是 [非减递增]
// 构造测试用例
// 对于 {3, 3, 3, 3, 3, 1, 3} 和 {3, 1,3, 3, 3, 3, 3} 程序没法判断向哪边收缩
// 程序判断向哪一遍收缩
