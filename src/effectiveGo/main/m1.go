package main
import (
	"fmt"
	"os"
	"effectiveGo/error"
)
func is()  {
	f, err := os.Open("t.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	d, err := f.Stat()
	if err != nil {
		f.Close()
		fmt.Println(err)
		return
	}
		fmt.Println(d)
	return
}


func Sqrt(f float64)(float64,error)  {
	if f<0{
		return 0,error1.New("math: square root of negative number")
	}
	return f,nil
}
//
//func main()  {
////	if 1>0{
////		fmt.Println(true)
////	}
////is()
//	f,err:=Sqrt(-1)
//	if err!=nil{
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(f)
//}
type myStruce struct {
	name bool
	userid int64
}
var structZero myStruce
var intZero int
var int32Zero int32
var int64Zero int64
var uintZero uint
var uint8Zero uint8
var uint32Zero uint32
var uint64Zero uint64
var byteZero byte
var boolZero bool
var float32Zero float32
var float64Zero float64
var stringZero string
var funcZero func(int)int
var byteArrayZero [5]byte
var boolArrayZero [5]bool
var byteSliceZero []byte
var boolSliceZero []bool
var mapZero map[string]bool
var interfaceZero interface{}
var chanZero chan int
var pointerZero *int
func main()  {
	fmt.Println("structZero: ", structZero)
	fmt.Println("intZero: ", intZero)
	fmt.Println("int32Zero: ", int32Zero)
	fmt.Println("int64Zero: ", int64Zero)
	fmt.Println("uintZero: ", uintZero)
	fmt.Println("uint8Zero: ", uint8Zero)
	fmt.Println("uint32Zero: ", uint32Zero)
	fmt.Println("uint64Zero: ", uint64Zero)
	fmt.Println("byteZero: ", byteZero)
	fmt.Println("boolZero: ", boolZero)
	fmt.Println("float32Zero: ", float32Zero)
	fmt.Println("float64Zero: ", float64Zero)
	fmt.Println("stringZero: ", stringZero)
	fmt.Println("funcZero: ", funcZero)
	fmt.Println("funcZero == nil?", funcZero == nil)
	fmt.Println("byteArrayZero: ", byteArrayZero)
	fmt.Println("boolArrayZero: ", boolArrayZero)
	fmt.Println("byteSliceZero: ", byteSliceZero)
	fmt.Println("byteSliceZero's len?", len(byteSliceZero))
	fmt.Println("byteSliceZero's cap?", cap(byteSliceZero))
	fmt.Println("byteSliceZero == nil?", byteSliceZero == nil)
	fmt.Println("boolSliceZero: ", boolSliceZero)
	fmt.Println("mapZero: ", mapZero)
	fmt.Println("mapZero's len?", len(mapZero))
	fmt.Println("mapZero == nil?", mapZero == nil)
	fmt.Println("interfaceZero: ", interfaceZero)
	fmt.Println("interfaceZero == nil?", interfaceZero == nil)
	fmt.Println("chanZero: ", chanZero)
	fmt.Println("chanZero == nil?", chanZero == nil)
	fmt.Println("pointerZero: ", pointerZero)
	fmt.Println("pointerZero == nil?", pointerZero == nil)
//is()
//	f,err:=Sqrt(-1)
//	if err!=nil{
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(f)
//	b()
	s:=[...][2]string{{"a","b"},{"c","d"}}
	b:=s
	for _,v:=range b{
		fmt.Println(v[1],v[0])
	}
}


func test()  {
	var p *[]int  =new ([]int)
//	var v []int=make([]int,100)
//	var p *[]int=new ([]int)
*p=make([]int,100,100)
}
func trace(s string)string  {
	fmt.Println("entering:",s)
	return s
}
func un(s string)  {
	fmt.Println("leaving:",s)
}
func a()  {
	defer un(trace("a"))
	fmt.Println("in a")
}
func b()  {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}