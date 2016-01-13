package heap
import "testing"


func Test_demo(t testing.T) {
	var bigHeap=NewBigHeap()
	array:=[]int{nil,1, 2, 5, 10, 3, 7, 11, 15, 17, 20, 9, 15, 8, 16}
	append(bigHeap,array...)
	bigHeap.Print()
}