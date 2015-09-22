package main
import (
	"fmt"
)
var isActive bool
var enabled, disable = true, false
func test() {
	var available bool
	valid := false
	available = true
	fmt.Print(available, valid, available)
}
var frenchHello string
var emptyString string = ""
func test1() {
	no, yes, maybe := "no", "yes", "maybe"
	japaneseHello := "Konichiwa"
	frenchHello = "Bonjour"
	m := `hello
	world`
	fmt.Print(m, no, yes, maybe, japaneseHello, frenchHello)
}
const (
	x = iota
	y = iota
	z = iota
	w
)
const v = iota
const (
	e, f, g = iota, iota, iota
)
var arr [10]int
func test2()  {
	arr[0]=42
	arr[1]=13
	fmt.Print(arr[0])
	fmt.Print(arr[9])

}
var slice=[]byte{'a','b','c','s'}


func main()  {
	str:=`hello
	world`
	fmt.Println(str)
}