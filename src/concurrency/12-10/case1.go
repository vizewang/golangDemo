package main


import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	fmt.Println(total)
	c <- total
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	//	go say("world")
	//	say("hello")
	//	a:=[]int{7,2,8,-9,4,0}
	//	c:=make(chan int)
	//	go sum(a[:len(a)/2],c)
	//	go sum(a[len(a)/2:],c)
	//	fmt.Println(<-c)
	//	fmt.Println(<-c)
	//	x,y:=<-c,<-c
	//	fmt.Println(x,y,x+y)

	//	c:=make(chan int,3)
	//	c<-1
	//	c<-2
	//	fmt.Println(<-c)
	//	fmt.Println(<-c)
	//	c := make(chan int)
	//	quit := make(chan int)
	//	go func() {
	//		for i := 0; i < 10; i++ {
	//			fmt.Println(<-c)
	//		}
	//		quit <- 0
	//	}()
	//	fibonacci2(c, quit)
	//	go fibonacci(cap(c),c)
	//	for i:=range c{
	//		fmt.Println(i)
	//	}
	fmt.Println(runtime.NumCPU())
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				o <- true
				break

			}
		}
	}()
	<-o
	fmt.Println("end")
}