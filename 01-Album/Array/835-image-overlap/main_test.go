package main 
// func largestOverlap(img1 [][]int, img2 [][]int) int {
	// Note also that a translation does not include any kind of rotation.?
	// 怎么去表达一个图片的 1 的模式，pattern ?
	// 
	// 
	// In addition, it could be a practical problem in real world. For     // instance, if one can find the maximal overlapping zone between  // two images, one could clip the images to make them smaller and // // more focused
	   
	// 代数知识
	// some knowledge of linear algebra (or geometry),   
	// solve it with the conception of convolution 卷积？
	//    
//   }

// 四个循环的暴力解法
// 能否用视觉化的效果表达相关概念？


type tuple struct {
    x int
    y int
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func largestOverlap(A [][]int, B [][]int) int {
    var pA []tuple
    var pB []tuple
    for i := 0; i < len(A); i++ {
        for j := 0; j < len(A[0]); j++ {
            if A[i][j] == 1 {
                pA = append(pA, tuple{i, j})
            }
            if B[i][j] == 1 {
                pB = append(pB, tuple{i, j})
            }
        }
    }
    
    mp := make(map[tuple]int)
    ans := 0
    for  _, a := range pA {
        for _, b := range pB {
            diff := tuple{a.x - b.x, a.y - b.y}
            mp[diff]++
            ans = max(ans, mp[diff])
        }
    }
    
    return ans
}

