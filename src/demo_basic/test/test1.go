package test
import "fmt"
func Test1() {
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a, b []byte
	a = ar[2:5]
	b = ar[3:5]
	fmt.Println(a[0])
	fmt.Println(b)
}