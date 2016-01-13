package main
import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var b *[]int
	b=&a[0]
	fmt.Println(b)
	fmt.Println(&a[0])
	s := a[0:]
	z:=&s[0]
	y:=*z
	fmt.Println(y)
	fmt.Println(&s[0])
	s=append(s, 11, 22, 33)
	fmt.Println(&s[0])
	sa := a[2:7]
	sb := sa[3:5]
	fmt.Println(a[0])
	fmt.Println(a, len(a), cap(a))    //输出：[1 2 3 4 5 6 7 8 9 0] 10 10
	fmt.Println(s, len(s), cap(s))    //输出：[1 2 3 4 5 6 7 8 9 0 11 22 33] 13 20
	fmt.Println(sa, len(sa), cap(sa)) //输出：[3 4 5 6 7] 5 8
	fmt.Println(sb, len(sb), cap(sb)) //输出：[6 7] 2 5
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{7, 8, 9}
	at := copy(s2, s1)
	fmt.Println(s2)
	fmt.Println(at)

	sc:=[]int{11,22,33,44,55,66}
//	i:=2
	sc=append(sc[:len(sc)-1])
	fmt.Println(sc)
}
