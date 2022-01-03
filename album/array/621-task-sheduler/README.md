### 题解
    // least number of units of times 
    // 
    
    // 题目特点： at least 
    // 动态规划、贪心
    
    // map["c"] = int x 
    // 数据结构： 双循环链表, 带 count 。 需要防止死循环。
    // link 循环，带 count  
    // 1. 怎么判断是否应该 idle ?
    //     - 回溯 n 个元素，判断是否与当前相等。 
    //     - 不能回溯时，可以直接放置。 
    // 2. ++ 1 的条件
    
    // 没想到的点
    // 1. 优先级队列
    // 这道题最关键的点是每次选取task的时候，这个task应该是当前可选择的task里面frequency最高的。我觉得如果要把task顺序输出出来的话就需要用到priority       // queue，这题因为只要输出一个长度所以有这个讨巧的解法，但反而让解题思路变的有些晦涩。


### 题目思路
1. My intuitive thinking of why this greedy works: If you choose the most frequent element, then you will have small probability of "collision".

2. 




### 题目解法一  pq 

来源于这个代码：https://leetcode.com/problems/task-scheduler/discuss/104501/Java-PriorityQueue-solution-Similar-problem-Rearrange-string-K-distance-apart

```
public int leastInterval(char[] tasks, int n) {
     Map<Character, Integer> map = new HashMap<>();
    for (int i = 0; i < tasks.length; i++) {
        map.put(tasks[i], map.getOrDefault(tasks[i], 0) + 1); // map key is TaskName, and value is number of times to be executed.
    }
    PriorityQueue<Map.Entry<Character, Integer>> q = new PriorityQueue<>( //frequency sort
            (a,b) -> a.getValue() != b.getValue() ? b.getValue() - a.getValue() : a.getKey() - b.getKey());

    q.addAll(map.entrySet());

    int count = 0;
    while (!q.isEmpty()) {
        int k = n + 1;
        List<Map.Entry> tempList = new ArrayList<>();
        while (k > 0 && !q.isEmpty()) {
            Map.Entry<Character, Integer> top = q.poll(); // most frequency task
            top.setValue(top.getValue() - 1); // decrease frequency, meaning it got executed
            tempList.add(top); // collect task to add back to queue
            k--;
            count++; //successfully executed task
        }

        // 优化点一
        for (Map.Entry<Character, Integer> e : tempList) {
            if (e.getValue() > 0) q.add(e); // add valid tasks 
        }

        if (q.isEmpty()) break;
        count = count + k; // if k > 0, then it means we need to be idle
    }
    return count;
}
```

- 可以利用一个 next 函数，将 `优化点一` 的 for 循环优化掉。 
### 解法二 

1. https://leetcode.com/problems/task-scheduler/discuss/146630/Golang-heap-solution

2. https://www.youtube.com/watch?v=YCD_iYxyXoo&t=340s


```
type Node struct {
    nextCycle, count int
}

type Nodes []*Node

func(n *Nodes) Push(v interface{}) {
    *n = append(*n, v.(*Node))
}

func(n *Nodes) Pop() interface{} {
    rtn := (*n)[len(*n)-1]
    *n = (*n)[:len(*n)-1]
    return rtn
}

func(n *Nodes) Swap(i,j int) {
    (*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}

func(n Nodes) Len() int {
    return len(n)
}

func(n Nodes) Less(i, j int) bool {
    return n[i].nextCycle < n[j].nextCycle || n[i].nextCycle == n[j].nextCycle && n[i].count > n[j].count    
}

func leastInterval(tasks []byte, n int) int {
    m := make(map[byte]*Node)
    
    for _, t := range tasks {
        if m[t] == nil {
            m[t] = &Node{0, 1}
        } else {
            m[t].count++
        }
    }
    
    nodes := Nodes{}
    for _, task := range m {
        nodes = append(nodes, task)
    }
    
    heap.Init(&nodes)
    
    cycle := 0
    
    for len(nodes) > 0 {    
        if nodes[0].nextCycle <= cycle {
            nodes[0].count--
            if nodes[0].count == 0 {
                heap.Pop(&nodes)
            } else {
                nodes[0].nextCycle += n + 1
                heap.Fix(&nodes, 0)
            }
        }
        cycle++
    }
    
    return cycle
}
```