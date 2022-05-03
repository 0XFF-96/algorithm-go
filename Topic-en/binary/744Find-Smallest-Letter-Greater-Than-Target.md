
### 这道题目本质上，是一道二分查找寻找 【 右边界 】 的题目


```

// func nextGreatestLetter(letters []byte, target byte) byte {
//     // Let's find the rightmost position to insert target into letters
//     // 
// }


// func nextGreatestLetter(letters []byte, target byte) byte {
//     low := 0
//     high := len(letters) - 1
//     mid := 0
    
//     //  is sorted in non-decreasing order
//     //  the smallest character in the array that is larger than target
//     //  Let's find the rightmost position to insert target into letters
//     for low < high {
//         mid = low + (high - low) / 2
        
//         if letters[mid] <= target {
//             low = mid + 1
//         } else {
//             high = mid - 1
//         }
//     }
    
//     // return letters[low] // out of range ?
//     // return letters[(low) % len(letters)]
// }

func nextGreatestLetter(letters []byte, target byte) byte {
    low, high := 0, len(letters)-1
    
    for low < high { // low == high will break 
        mid := (low + high)/2
        midVal := letters[mid]
        
        if midVal <= target {
            low = mid + 1
        } else {
            high = mid  // remains
        }
    }
    
    if letters[low] > target {
        return letters[low]
    } else {
        return letters[0]
    }
}

```
