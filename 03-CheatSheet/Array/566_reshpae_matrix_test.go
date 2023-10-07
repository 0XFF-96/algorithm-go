package array


// Do you know how 2d matrix is stored in 1d memory? Try to map 2-dimensions into one.
//
//
// M[i][j]=M[n*i+j] , where n is the number of cols. This is the one way of converting 2-d indices into one 1-d index. Now, how will you convert 1-d index into 2-d indices?
// Try to use division and modulus to convert 1-d index into 2-d indices.
// M[i] => M[n/i][n%i] Will it result in right mapping? Take some example and check this formulae.

// 题目类型： array 下标的相关运算
// 解题技巧
// know it, use it, and then improve it
//