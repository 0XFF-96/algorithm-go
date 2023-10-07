# 如何设计一个哈希表

### 场景

- 散列表的装载因子 = 填入表中的元素个数 / 散列表的长度

- https://learnku.com/articles/45870
- https://segmentfault.com/a/1190000039219752


### 自然语言描述

哈希表存储的是由键（key）和值（value）组成的数据。
在哈希表中，我们可以利用哈希函数快速访问到数组中的目标数据。

如果发生哈希冲突(哈希函数对不同的键, 返回了相同的哈希值)，
就使用链表进行存储(因此, 哈希表一般至少是数组+链表的二维结构)。

如果数组的空间太小，
使用哈希表的时候就容易发生冲突，
线性查找的使用频率也会更高；
反过来，如果数组的空间太大，
就会出现很多空箱子，造成内存的浪费。

### 目标

实现一个哈希表, 提供对键值对数据的增删改查, 且性能不能太差
物理结构

采用分区 + 数组 + 链表的三维结构
分区是一个双向链表, 由小到大呈2次幂增长
当前分区总是指向最新也是最大的那个分区
查找某个键时, 需要遍历所有分区
哈希函数

合适的哈希函数对性能影响非常巨大
对哈希函数的要求一般是足够快, 足够"散", 对不同键值最好能随机均匀分布
本示例采用hash/crc64, 多项式为ECMA
如果哈希函数的计算比较耗时, 则应当尽可能重复利用计算结果
扩容与缩容

扩容

填充因子为3/4, 即当前分区的元素数>容量的3/4时, 创建2倍大的新分区
扩容不做任何拷贝和rehash.
扩容后, 非当前分区变成只读和只删.
缩容: 当某个分区未持有任何元素, 且不为当前分区时, 删除此分区

### 核心代码

const (
    expandFactor = 0.75
    )

type HashMap struct{
    m []*KeyPairs //hash表的桶
    cap int
    len int
}


//写入一个键值对
func (h *HashMap)Put(key string, value interface{}) {
    index := h.Index(key)
    element := h.m[index]
    if element == nil { //该下标没有值，直接写入
        h.m[index] = &KeyPairs{key,value, nil}
    }else{ //有值，找到最后一个节点
        for element.next != nil {
            if element.key == key { //如果是相同的键就进行修改操作
                element.value = value
                return
            }
            element = element.next
        }
        element.next = &KeyPairs{key,value,nil}
    }
    h.len++
    if (float64(h.len)/float64(h.cap))>expandFactor { //需要扩容，把切片容量变为2倍
        newH :=  NewHashMap(2*h.cap)
        for _,pairs := range h.m {
            for pairs != nil {
                newH.Put(pairs.key,pairs.value)
            }
        }
        h.cap = newH.cap
        h.m = newH.m
    }
}
