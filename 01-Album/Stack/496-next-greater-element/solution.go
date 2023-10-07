package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// next greater element 的定义
	// brute force 的解法
	// map index
	m := map[int]int {}
	for idx, n := range nums2 {
		m[n] = idx
	}

	var ret []int
	for _, n := range nums1 {
		idx := m[n]

		found := false
		for i:=idx+1; i < len(nums2); i++{
			if nums2[i] > n {
				found = true
				ret = append(ret, nums2[i])
				break
			}
		}
		if !found {
			ret = append(ret, -1)
		}
	}
	return ret
}


// 
func nextGreaterElementV2(nums1 []int, nums2 []int) []int {
    var result []int
    incrementMap := make(map[int]int,len(nums2))
    
    for i, num2 := range nums2 {
        var found bool
        for _, greaterValue := range nums2[i+1:] {
            if greaterValue > num2 {
			 // next incremental number
                incrementMap[num2] = greaterValue
                found =true
                break
            }
        }
        if found == false {
            incrementMap[num2] = -1
        }
    } 
    for _, num1 := range nums1 {
        result = append(result, incrementMap[num1])
    }
    
    return result
}

func nextGreaterElementV3(nums1 []int, nums2 []int) []int {
    stack := make([]int, len(nums2))
	m := map[int]int{}
	
	// 构建 monitic stack
    for i, v := range nums2 {
        for len(stack)>0 && nums2[stack[len(stack)-1]] < v {
            m[nums2[stack[len(stack)-1]]] = v
            stack = stack[0:len(stack)-1]
        }
        stack = append(stack, i)
    }

    for len(stack)>0 {
        m[nums2[stack[len(stack)-1]]] = -1
        stack = stack[0:len(stack)-1]
    }

    res := make([]int, len(nums1))
    for i, v := range nums1 {
        res[i] = m[v]
    }
    return res
}