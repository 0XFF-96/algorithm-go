// 参考相关文档
// 相关: Flatten Nested List Iterator
// https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/
// 


/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

 type NestedIterator struct {
    nums []int
    p int
    
    // 优化点1: 不需要一开始就把遍历的元素
    // 有一些可能之后，并不需要这个元素的。
}

// NOTE: https://leetcode.com/problems/flatten-nested-list-iterator/discuss/571564/Golang-4ms
// 

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    // construct list from nestedInteger
    // use stack for holding a un traverse element
    
    var res NestedIterator
    var flatten func([]*NestedInteger)
    flatten = func(nestedList []*NestedInteger){
        for _, nest :=  range nestedList{
            if nest.IsInteger() {
                res.nums = append(res.nums, nest.GetInteger())
            } else {
                flatten(nest.GetList())
            }
        }
    }
    flatten(nestedList)
    res.p = 0
    return &res
}

func (this *NestedIterator) Next() int {
    this.p++
    return this.nums[this.p - 1]
}

func (this *NestedIterator) HasNext() bool {
    return this.p   < len(this.nums) 
}

