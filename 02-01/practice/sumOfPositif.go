
package main

import "fmt"

func PositiveSum(numbers []int) int {
var x int
for _, num := range numbers {
if num > 0 {
x += num
}
}
return x
}

func main() {
numbers := []int{1, 2, 3, -4, -5, 6}
fmt.Println(PositiveSum(numbers))
}


