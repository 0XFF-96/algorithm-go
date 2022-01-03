package main

func findPoisonedDuration(timeSeries []int, duration int) int {
    if len(timeSeries) == 0 {
        return 0 
    }
    
    if len(timeSeries) == 1 {
        return duration
    }
    // in ascending time 
    count := 0 
    for i:=1; i < len(timeSeries); i ++ {
        elaspTime := timeSeries[i] - timeSeries[i-1]
        if elaspTime > duration {
            count += duration
        } else {
            count += elaspTime
        }
    }
    // add edge condition
    count += duration
    return count
}