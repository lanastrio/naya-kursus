
package main

import (
"fmt"
"strings"
)

func ReverseWords(str string) (result string) {
strs := strings.Fields(str)
for _, value := range strs {
result = value + " " + result
}
return result[:len(result)-1]
}

func main() {
fmt.Println(ReverseWords("hello world!"))
fmt.Println(ReverseWords("yoda doesn't speak like this"))
fmt.Println(ReverseWords("foobar"))
fmt.Println(ReverseWords("kata editor"))
fmt.Println(ReverseWords("row row row your boat"))

}
