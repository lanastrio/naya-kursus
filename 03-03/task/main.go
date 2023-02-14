package main
import
"fmt"
 func plusOne(digits []int) []int {
    lastElemIndex := len(digits) - 1
    for i := lastElemIndex; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        digits[i] = 0
    }
    return append([]int{1}, digits...)
}

func main(){
	digits := [] int {1,9} 
	fmt.Println(plusOne(digits))
}
