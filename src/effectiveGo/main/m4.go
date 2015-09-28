package main
import (
	"fmt"
	"runtime"
)
func main() {
//	c := make(chan int)
//	go func() {
//		time.Sleep(100000)
//		fmt.Println("h1")
//		c <- 1
//		fmt.Println("h3")
//	}()
////	fmt.Println(<-c)
//	time.Sleep(10000000)
//	fmt.Println("h2")
//
//	fmt.Println(runtime.NumCPU())
//	fmt.Println(runtime)
	v:=Vector{1,2,3,4,5,6,1,2,3,4,5,6}
	v.DoAll(v)
}


type Vector []float64
func (v Vector)Op( f float64)(float64)  {
	fmt.Println(f)
	return 22
}
func (v Vector)DoSome(i,n int,u Vector,c chan int)  {
	for ;i<n;i++{
		v[i]+=u.Op(v[i])
	}
	c<-0
}

func (v Vector)DoAll(u Vector){
	NCPU:=runtime.NumCPU();
	c:=make(chan int,NCPU)
	for i:=0;i<NCPU;i++{
		go v.DoSome(i*len(v)/NCPU,(i+1)*len(v)/NCPU,u,c)
	}
	for i:=0;i<NCPU;i++{
		fmt.Println(<-c)
	}
}