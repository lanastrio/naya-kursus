
package main

import "fmt"

func mergeNumbers(numbers []int) {
	
sum := numbers[0] + numbers[1]

if numbers[0]+numbers[1] <= 10 {
value := sum
// value := append(numbers, sum)
fmt.Println(value)
}
}

func main() {
mergeNumbers([]int{1, 2, 3, 4, 5})
}
