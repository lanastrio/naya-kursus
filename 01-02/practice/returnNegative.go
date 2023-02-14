package main
import

"fmt"

func MakeNegative(x int) int {
if x >= 0 {
return -x
}
return x
}
func main(){
x:=0
hasil:=MakeNegative(x)
fmt.Println(hasil)
}

