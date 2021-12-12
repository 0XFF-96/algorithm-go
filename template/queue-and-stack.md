
s1 s2 画图的摆放位置

s1            s2
[队头, , , , ]  [, , , 队尾]

class MyQueue {
    
    /** 添加元素到队尾 */
    public void push(int x);
    
    /** 删除队头的元素并返回 */
    public int pop();
    
    /** 返回队头元素 */
    public int peek();
    
    /** 判断队列是否为空 */
    public boolean empty();
}

/** 返回队头元素 */
public int peek() {
    if (s2.isEmpty())
        // 把 s1 元素压入 s2
        while (!s1.isEmpty())
            s2.push(s1.pop());
    return s2.peek();
}

/** 删除队头的元素并返回 */
public int pop() {
    // 先调用 peek 保证 s2 非空
    peek();
    return s2.pop();
}

/** 判断队列是否为空 */
public boolean empty() {
    return s1.isEmpty() && s2.isEmpty();
}

// 用队列实现栈是没啥亮点的问题，但是用双栈实现队列是值得学习的
// 