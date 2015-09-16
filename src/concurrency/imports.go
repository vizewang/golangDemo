package main
import (
	"fmt"
	"math"
)
func add(x,y int)int{
	return x+y
}
func main(){
	fmt.Println("Now you have %g problems.",math.Nextafter(2,3))
}
