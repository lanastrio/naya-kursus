package main
import 
"fmt"

func IsTriangle(a,b,c int)bool{ 
cok := (a + b) > c
	if !cok {
		return false
	}
	cok = (a + c) > b
	if !cok {
		return false
	}
	cok = (b + c) > a
	if !cok {
		return false
	}
	return true
}
func main(){
	fmt.Println(IsTriangle(5,1,2))
	fmt.Println(IsTriangle(5,1,5))
	fmt.Println(IsTriangle(2,2,2))
}