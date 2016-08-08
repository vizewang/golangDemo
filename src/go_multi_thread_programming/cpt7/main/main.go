package main

import (
	"fmt"
	"time"
	"runtime"
	//"go_multi_thread_programming/cpt7/others"
)

func main() {
	//Demo1()
	//Demo4()
	//others.PersonHandlerDemo()
	Demo6()
}

func Demo1() {
	fmt.Println("start");
	ch := make(chan int, 5)
	sign := make(chan byte, 2)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch)
		fmt.Println("The channel is closed.")
		sign <- 0
	}()
	go func() {
		for {
			e, ok := <-ch
			fmt.Printf("%d(%v)\n", e, ok)
			if !ok {
				break
			}
			time.Sleep(2 * time.Second)
		}
		fmt.Println("Done.")
		sign <- 1
	}()
	<-sign
	<-sign
}

func Demo2() {
	go println("Go!Goroutine")
	time.Sleep(time.Millisecond)
}

func Demo3() {
	name := "Eric"
	go func() {
		fmt.Printf("Hello,%s.\n", name)
	}()
	name = "Harry"
	runtime.Gosched()
}

func Demo4() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello,%s.\n", who)
		}(name)
	}
	runtime.Gosched()
}

var ch3 chan int
var ch4 chan int
var chs = []chan int{ch3, ch4}
var numbers = []int{1, 2, 3, 4, 5}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i];
}

func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)
	return chs[i]
}
func Demo5() {
	select {
	case getChan(0) <- getNumber(2):
		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(3):
		fmt.Println("2th case is selected.")
	default:
		fmt.Println("default")
	}
}

func Demo6() {
	chanCap:=5
	ch7:=make(chan  int,chanCap)
	for i:=0;i<chanCap;i++{
		select {
		case ch7<-1:
		case ch7<-2:
		case ch7<-3:
		}
	}
	for i:=0;i<chanCap;i++{
		fmt.Printf("%v\n",<-ch7)
	}
}