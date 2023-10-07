
import "math"

func asteroidCollision(asteroids []int) []int {
    var stack []int 
    for _, rock := range asteroids {
        collide := false
        for len(stack) != 0 && rock<0 && stack[len(stack)-1]>0 {
            if rock + stack[len(stack)-1] == 0 {
                collide = true
                stack = stack[:len(stack)-1]
                break
            } else if math.Abs(float64(rock)) > math.Abs(float64(stack[len(stack)-1])) {
                stack = stack[:len(stack)-1]
            } else {
                collide = true
                break
            }
        }
        if !collide {
            stack = append(stack,rock)
        }
    }
    return stack
}