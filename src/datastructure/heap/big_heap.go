package heap
import (
	"fmt"
)

type heapArray  interface {}

func NewBigHeap() heapArray {
	h:=make(interface{},0,10)
	return h
}

func (this *heapArray)Insert(value int) {
	if(len(*this)==0){
		*this=append(*this,0)
	}
	fmt.Println(this)
	*this=append(*this,value)
}

func heapUp(this *heapArray, index int) {
	if(index>1){
		parent:=index/2
		s:=*this
		parentValue:=s[parent]
		indexValue:=s[index]
		if(parentValue<indexValue){
			swap(this,parent,index)
			heapUp(this,parent)
		}
	}
}

func (this *heapArray)Delete(index int) {

	*this=append(*this[:index],*this[index+1:]...)

	*this=append(*this[:len(*this)-1])
}

func heapDown(this *heapArray,index int) {
	n:=len(this)-2
	child:=-1;
	if 2*index>n{
		return
	}else if 2*index<n {
		child=2*index
		if this[child]<this[child+1]{
			child++
		}
	}else if 2*index==n {
		child=2*index
	}

	if this[child]>this[index]{
		swap(this,child,index)
		heapDown(this,child)
	}
}

func (this *heapArray)Adjust() {
	for i:=len(this);i>0;i--{

	}
}

func (this *heapArray)Print() {
	fmt.Println(this)
}

func adjust(this *heapArray, i, n int) {
	var child int;
	for ;i<=n/2;{
		child=i*2;
		if child+1<=n&&this[child]<this[child+1]{
			child+=1;
		}
		if this[i]<this[child]{
			swap(this,i,child)
			i=child
		}
		break
	}
}

func (this *heapArray)HeapSort() {
	for i:=len(this)-1;i>0;i--{
		swap(this,1,i)
		adjust(this,1,i-1)
	}
}

func swap(target *heapArray,a,b int) {
	tmp:=target[a]
	target[a]=target[b]
	target[b]=tmp
}