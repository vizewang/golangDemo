package main
import "fmt"

type Speaker interface  {
	Say(string)
	Listen(string) string
	Interrupt(string)
}
type WangLan struct  {
	msg string
}
func (this *WangLan)Say(msg string) {
	fmt.Printf("王兰说:%s\n",msg)
	
}
func (this *WangLan) Listen(msg string)string {
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

type Html[]interface{}
func main() {
	wl:=&WangLan{}
	jl:=&JiangLou{}
	var person Speaker
	person=wl
	person.Say("Hello world!")
	person=jl
	person.Say("Good luck!")


	fmt.Println("-----------------------------demo2")
	html:=make(Html,5)
	html[0] = "div"
	html[1] = "span"
	html[2] = []byte("script")
	html[3] = "style"
	html[4] = "head"
	for index,element:=range html{
		if value,ok:=element.(string);ok{
			fmt.Printf("html[%d] is a string and its value is %s \n",index,value)
		}else if value,ok:=element.([]byte);ok{
			fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
		}
	}
}