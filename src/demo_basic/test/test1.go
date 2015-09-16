package test
import "fmt"
func Test1() {
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a, b []byte
	a = ar[2:5]
	b = ar[3:5]
	fmt.Println(a[0])
	fmt.Println(b)
}

var numbers map[string]int
func Test2()  {
	numbers:=make(map[string]int)
	numbers["one"]=1
	numbers["ten"]=10
	numbers["three"]=3
	fmt.Println("第三个数字是：",numbers["three"])
	delete(numbers,"three")
	fmt.Println(numbers)
}
func Test3(x int)  {
	if x>10{
		fmt.Println("x is greater than 10")
	}else{
		fmt.Println("x it less than 10")
	}

}
func Test4()  {
	i:=0
	Here:
	fmt.Println(i)
	i++
	if(i<2){
		goto Here
	}

}
func Test5()  {
	sum:=0
	for index:=0;index<10;index++{
		sum+=index
	}
	fmt.Println("sum is equal to ",sum)
}

func Test6()  {
	i:=10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2,3,4:
		fmt.Println("i is equal to 2,3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
		fallthrough
	default:
		fmt.Println("All I know is what i is an integer")

	}
	
}