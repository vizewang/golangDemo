package main
import "fmt"

func main() {
	b:=new(bool)
	fmt.Println(*b)
	i:=new(int)
	fmt.Println(*i)
	s:=new(string)
	fmt.Println(*s)

	p:=new([]int)
	*p=make([]int,100)

	fmt.Println(*p)

}


