package main
import (
	"reflect"
	"fmt"
)

type T struct {
	A int
	B string
}

func main() {
	var x float64=3.4
	v:=reflect.ValueOf(&x)
	w:=v.Elem()
	fmt.Println("type or v:",w.Type())
	fmt.Println(w.CanSet())
	t:=T{23,"skidoo"}
	s:=reflect.ValueOf(&t).Elem()
	typeOfT:=s.Type()
	for i:=0;i<s.NumField();i++{
		f:=s.Field(i)
		fmt.Printf("%d:%s %s=%v\n",i,typeOfT.Field(i).Name,f.Type(),f.Interface())
	}

//	var val interface{}=nil
//	if val==nil{
//		fmt.Println("val is nil")
//	}else {
//		fmt.Println("val is not nil")
//	}

	var val interface{}="hello"
	fmt.Println(val)
	val=[]byte{'a','b','c'}
	fmt.Println(val)
//	v.SetFloat(7.1)
//	fmt.Println("type:",v.Type())
//	fmt.Println("kind is float64:",v.Kind()==reflect.Float64)
//	fmt.Println("value:",v.Float())
//	var y uint8='x'
//	w:=reflect.ValueOf(y)
//	fmt.Println("type:",w.Type())
//	fmt.Println("kind is uint8:",w.Kind()==reflect.Uint8)
//	y=uint8(w.Uint())
//	y=w.Uint()
}
