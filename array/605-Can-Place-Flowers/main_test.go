package main 

func canPlaceFlowers(flowerbed []int, n int) bool {
    i := 0 
    count := 0 
    
    for i < len(flowerbed) {
        // 这个判断有点难看
        // 简化一下
        if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed) -1 || flowerbed[i+1] == 0) {
            
            // 这里是先放置，后增加。
            flowerbed[i] = 1
            i++
            count++
        }
        
        if count >= n {
            return true 
        }
        i++
    }
    return false 
}


