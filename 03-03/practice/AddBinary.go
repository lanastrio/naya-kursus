package main

import
"fmt"
func addBinary(a string, b string) string {
	ar := []byte(a)
	br := []byte(b)
	sum := ""

	i := len(ar) - 1
	j := len(br) - 1

	currSum := byte(0)
	add := byte(0)

	for i >= 0 || j >= 0 || add > 0 {
		
		dar := byte(0)
		dbr := byte(0)

		if i >= 0 {
			dar = ar[i] - '0'
		}
		if j >= 0 {
			dbr = br[j] - '0'
		}
		currSum = dar + dbr + add

		add = currSum >> 1 
		
		sum = string(currSum%2+'0') + sum

		i--
		j--

		
	}

	return sum
}
func main(){
	a:= "11"
	b:= "1"
	fmt.Println(addBinary(a,b))
}