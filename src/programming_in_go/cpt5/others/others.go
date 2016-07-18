package others

import (
	"fmt"
	"math/rand"
	"math"
)

func Demo1() {
	var i interface{} = 99
	var s interface{} = []string{"left", "right"}
	j := i.(int)
	fmt.Printf("%T->%d\n", i, j)
	if i, ok := i.(int); ok {
		fmt.Printf("%T->%d\n", i, j)
	}
	if s, ok := s.([]string); ok {
		fmt.Printf("%T->%q\n", s, s)
	}
}
func Demo2() {
	classifier(5, -17.9, "ZIP", nil, true, complex(1, 1))
}
func Demo3() {
	counterA:=createCounter(2)
	counterB:=createCounter(102)
	for i:=0;i<5;i++{
		a:=<-counterA
		fmt.Printf("(A->%d,B->%d)",a,<-counterB)
	}
	fmt.Println();
}
func Demo4() {
	makeChannel();
}
func classifier(items ... interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is an int\n", i)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("param #%d is an unsigned int\n", i)
		case nil:
			fmt.Printf("param #%d is a nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d is a unknow\n", i)
		}
	}
}

func createCounter(start int) chan int {
	next := make(chan int)
	go func(i int) {
		for {
			next <- i
			i++
		}
	}(start)
	return next
}

func makeChannel()  {
	channels:=make([]chan bool,6)
	for i:=range channels{
		channels[i]=make(chan bool)
	}
	go func() {
		for{
			channels[rand.Intn(6)]<-true
		}
	}()

	for i:=0;i<36;i++{
		var x int
		select {
		case <-channels[0]:
			x=1
		case <-channels[1]:
			x=2
		case <-channels[2]:
			x=3
		case <-channels[3]:
			x=4
		case <-channels[4]:
			x=5
		case <-channels[5]:
			x=6
		}

		fmt.Printf("%d",x)
		fmt.Println();
	}
}

func ConvertInt64ToInt(x int64) int {
	if math.MinInt32 <=x && x<=math.MaxInt32{
		return int(x)
	}
	panic(fmt.Sprintf("%d is out of the int32 range",x))
}
func IntFromInt64(x int64) (i int,err error) {
	defer func() {
		if e:=recover();e!=nil{
			err=fmt.Errorf("%v",e)
		}
	}()
	i=ConvertInt64ToInt(x)
	return i,nil;

}