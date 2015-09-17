package multiThread
import (
	"runtime"
	"fmt"
)


func Say(s string)  {
	for i:=0;i<5;i++{
		runtime.Gosched()
		fmt.Println(s)
	}
}

func Main1()  {
//	go Say("world")
	Say("hello")
}

