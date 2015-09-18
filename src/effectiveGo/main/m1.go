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
	f,err:=Sqrt(-1)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(f)
}