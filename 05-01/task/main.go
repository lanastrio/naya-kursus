package main
import

"fmt"

func Number(busStops [][2]int) int {
	// fmt.Println(busStops)
	var naik int
	var turun int
	for _, array := range busStops{
		naik+=array[0]
		turun+=array[1]
	}

	return naik-turun
}

func main(){
	x := [][2]int{{10,0},{3,5},{5,8}}
	Number(x)
	 fmt.Println(Number(x))
}