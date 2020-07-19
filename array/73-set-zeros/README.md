1- https://leetcode.com/problems/set-matrix-zeroes/
2- 参考题目难点。

### 难点
1. 这道题有不同的解法，都很值得参考。尤其是对于空间的优化来说。 
2. 优化空间的通常套路是，例如题目中给出的【数组】，看能不能提前借用一些空间。
3. 这道题目通常的犯错点是，没有 generation 的概念。有两种 0 ， 一种是原来数组就有的 0 ，另一种是后来才被设置为 0 的零。 
4. https://leetcode.com/problems/set-matrix-zeroes/discuss/26032/Golang-concise-solution-with-explanation