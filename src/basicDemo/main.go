package main
import (
	"fmt"
	"errors"
	"time"
)

func demo1() {
	m:=`c

	d`
	fmt.Println(m)
}

func demo2() {
	err:=errors.New("test error")
	if err!=nil{
		fmt.Println(err)
	}
}

func demo3(s interface{}) {
	a:=s.([2]interface{})

	tmp:=a[0]
	a[0]=a[1]
	a[1]=tmp
	fmt.Println("func slice ",a)
}
type sl []interface{}
func (this *sl)insert() {
	*this=append(*this,1)
	tmp:=*this
	fmt.Println(tmp[0])
	tmp[0]=2
}
type  testInt func(int) bool
func isOdd(integer int) bool {
	if integer%2==0{
		return false
	}
	return true
}
func isEven(integer int) bool {
	if integer%2==0{
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _,value:=range slice{
		if f(value){
			result=append(result,value)
		}
	}
	return result
}
func New()sl {
	return sl{}
}

func demo4() {
	t:=time.Date(2009,time.November,10,23,0,0,0,time.UTC)
	fmt.Println("go launched at %s\n",t.Local())
}

func demo5() {
	fmt.Println(32<<20)
}
func main() {
//	demo2()
//	s:=[2]interface{}{1,2}
//	demo3(s)
//	fmt.Println(s)
//	var s1 sl
//	s1.insert()
//	fmt.Println(s1)
//
//	var s2=New()
//	s2.insert()
//	fmt.Println(s2)
//
//	slice := []int {1, 2, 3, 4, 5, 7}
//	fmt.Println("slice = ", slice)
//	odd := filter(slice, isOdd)    // 函数当做值来传递了
//	fmt.Println("Odd elements of slice are: ", odd)
//	even := filter(slice, isEven)  // 函数当做值来传递了
//	fmt.Println("Even elements of slice are: ", even)
//
//	var s3=make([]interface{},2)
//	s3=append(s3,1,2)
//	fmt.Println(s3[2])
	demo5()
}