package main
import
"fmt"

func containsDuplicate(nums []int) bool {

	x:= make(map[int]int)

	for _, y:= range nums {

		x[y]++  
		if x[y] > 1 {
			return true
		}
	}
	return false
}

func main(){
	hasil := [] int {1,2,3,4}
	fmt.Println(containsDuplicate(hasil))
}