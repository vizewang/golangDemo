package main
import (
	"fmt"
)

func sum(a[]int,c chan int)  {
	total:=0
	for _,v:=range a{
		total+=v
	}
	c<-total
}

func main() {
//	test.Test1()
//	test.Test2()
//	test.Test3(11)
//	test.Test4()
//	test.Test5()
//	test.Test6()
//	test.Main()
//	test.Main3()
//	test.Main4()
//	test.Main5()
//	fmt.Println(test.ThrowsPanic(test.Main5))
//	test.Main9()
//	test.Main7()
//	test.Main10()
//	test1.Main1()
//	interfaces.Main3()
//	interfaces.Main4()
//	go multiThread.Say("world")
//	multiThread.Main1()

//	a:=[]int{7,2,8,-9,4,0}
//	c:=make(chan int)
//	go sum(a[:len(a)/2],c)
//	go sum(a[len(a)/2:],c)
//	go sum(a[len(a)/3:],c)
//	x,y:=<-c,<-c
//	z:/=<-c
//	fmt.Println(z)
//	fmt.Println(x,y,x+y)


	c:=make(chan int,1)
	c<-1
	c<-2
	fmt.Println(<-c)
	fmt.Println(<-c)

}


