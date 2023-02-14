package main
import (
"fmt"
)
func NbYear(p0 int, percent float64, aug int, p int) int {
  var c int
  sum:=p0
  for sum < p{
         sum = 1 + int(float64(sum) * (percent/100))+ aug
  c++
}
return c
}

func main(){
	fmt.Println(NbYear(1500, 5, 100, 5000))
	fmt.Println(NbYear(1500000, 2.5, 10000, 2000000))
	fmt.Println(NbYear(1500000, 0.25, 1000, 2000000))
	fmt.Println(NbYear(1500000, 0.25, -1000, 2000000))
}