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
func main()  {
//	if 1>0{
//		fmt.Println(true)
//	}
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