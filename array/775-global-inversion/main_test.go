package main 

import "math"
func isIdealPermutation(A []int) bool {
    if len(A) == 0 {
        return false 
    }
    // A = [1, 0, 2]
    //  global inversion, 1 和 0 ， local inverion , 1 和 0 
    
    // A = [1 , 2, 0 ]
    // global inversion [1,2] [1,0]  local inversion [2, 0]
    // 所以不符合
    
    for i := 0; i < len(A); i++ {
        a :=  math.Abs(float64(i - A[i]))
        if int(a) > 1 {
            return false
        }
    }
    return true 
}

// 大概可以用 merge sort 解决
// https://www.youtube.com/watch?v=3QHSJSFm0W0
//    public boolean isIdealPermutation(int[] A) {

//         for (int i = 0; i < A.length; i++) {
//             if (Math.abs(i - A[i]) > 1)
//                 return false;
//         }

//         return true;
//     }