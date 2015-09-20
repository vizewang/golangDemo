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


//	c:=make(chan int,1)
//	c<-1
//	c<-2
//	fmt.Println(<-c)
//	fmt.Println(<-c)

//	c:=make(chan int,10)
//	go fibonacci(cap(c),c)
//	for i:=range c{
//		fmt.Println(i)
//	}

	c:=make(chan int)
	quit:=make(chan int)
	go func() {
	for i:=0;i<10;i++{
		fmt.Println(<-c)
	}
	quit<-0
	}()
	fibonacci2(c,quit)
}

func fibonacci(n int,c chan int){
	x,y:=1,1
	for i:=0;i<n;i++{
		c<-x
		x,y=y,x+y
	}
	close(c)
}

func fibonacci2(c,quit chan int)  {
	x,y:=1,1
	for{
		select {
		case c<-x:
			x,y=y,x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}