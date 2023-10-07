package main 

func isOneBitCharacter(bits []int) bool {
    // Keep track of where the next character starts. At the end, you want to know if you started on the last bit
    // 
    
    // No idea what this quesiton is asking for.....
    //  Increment Pointer
    // 字符串一定会在 0 结束？
    // 
    // if and only if there are an even number of ones present.
    // 
        i := 0 
    
    // 为什么需要减去 1 
    for i < len(bits) -1{
        i += bits[i] + 1
    }
    return i == len(bits) -1
}