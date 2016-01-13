package array_list
import (
	"testing"
	"fmt"
)

func Test_Main(t *testing.T) {
	list:=NewList();
	fmt.Println(list)
	if MAXSIZE==len(list.Element){
		t.Log("right")
	}else {
		t.Log(len(list.Element))
	}
}
