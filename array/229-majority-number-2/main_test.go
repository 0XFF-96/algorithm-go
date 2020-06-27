func TestMajorityElement(t *testing.T){
	// to-do
}

func majorityElement(nums []int) []int {
    count1, count2, n1, n2 := 0, 0, 0, 0
    
	// 1. Calcualting counters and finding candidates n1 and n2
    for _, a := range nums {
        if a == n1 {
            count1++
        } else if a == n2 {
            count2++
        } else if count1 == 0 {
            n1 = a
            count1 = 1
        } else if count2 == 0 {
            n2 = a
            count2 = 1
        } else {
            count1--
            count2--
        }
    }
    
	// 2. Checking counters and forming output array
    var out []int
    count1 = len(nums)/3
    count2 = len(nums)/3
    for _, a := range nums {
        if a == n1 {
            count1--
        } else if a == n2 {
            count2--
        }   
    }
        
    if count1 < 0 {
        out = append(out, n1)
    }
        
    if count2 < 0 {
        out = append(out, n2)
    }
    
    return out
}