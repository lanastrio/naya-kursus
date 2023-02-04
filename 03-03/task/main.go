package main
import
"fmt"
 func plusOne(digits []int) []int {
    lastElemIndex := len(digits) - 1
	// Reading from last to first
    for i := lastElemIndex; i >= 0; i-- {
		// If less than 9 then just + 1 and done
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
		// otherwise set to 0 because +1 to next index
        digits[i] = 0
    }
	// return the new array
	// []int{1} is for integers like 9999, so we move 1 to the front -> 10000
    return append([]int{1}, digits...)
}
func main(){
	digits := [] int {4,3,2,1}
	fmt.Println(plusOne(digits))
}
