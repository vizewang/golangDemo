package main
import "fmt"



//func main()  {
////	var gid int32=0x12345678
////	var uid int8  =int8(gid)
////	fmt.Printf("uid=%#x,gid=%#x\n",uid,gid)
////var val interface{}="hello"
////	fmt.Println(val)
////	val=[]byte{'a','b','c'}
////	fmt.Println(val)
//
//	wl:=&WangLan{}
//	jl:=&JiangLou{}
//	fmt.Println(wl.msg,jl.msg)
//	var person Speaker
//	person=wl
//	person.Say("Hello World!")
//	person = jl
//	person.Say("Good Luck!")
//}

type Speaker interface {
	Say(string)
	Listen(string)string
	Interrupt(string)
}

type WangLan struct {
	msg string
}
func(this *WangLan)Say(msg string){
	fmt.Printf("王兰说:%s\n",msg)
}
func (this *WangLan)Listen(msg string)string  {
	this.msg=msg
	return msg
}

func (this *WangLan) Interrupt(msg string) {
	this.Say(msg)
}

type JiangLou struct {
	msg string
}

func (this *JiangLou) Say(msg string) {
	fmt.Printf("江娄说：%s\n", msg)
}

func (this *JiangLou) Listen(msg string) string {
	this.msg = msg
	return msg
}

func (this *JiangLou) Interrupt(msg string) {
	this.Say(msg)
}