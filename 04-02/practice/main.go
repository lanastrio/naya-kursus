package main
import(
"fmt"
)
func addDigits(num int) int {
	
    if num == 0 {
        return 0
    }
    if num % 9 == 0 {
        return 9
    }
    return num%9


}
func main(){
	fmt.Println(addDigits(28))
	fmt.Println(addDigits(0))
	fmt.Println(addDigits(9))
	fmt.Println(addDigits(1))
}

