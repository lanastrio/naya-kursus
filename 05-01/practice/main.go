package main
import 
"fmt"
func solution(str, ending string) bool {
	if len(str) < len(ending) {
	  return false
	}
	if str[len(str)-len(ending):] == ending {
		return true
	}
	return false
  }
  func main(){
	  fmt.Println(solution("abc","bc"))
	  fmt.Println(solution("abc","d"))
	  fmt.Println(solution("b","abc"))
  }