package array_list

const MAXSIZE = 20
type List struct {
	Element [MAXSIZE]int;
	length  int
}

func NewList() *List {
	return &List{Element:[MAXSIZE]int{}, length:0,}
}
func (I *List)InitList(d int, p int) {
	I.Element[p]=d
	I.length++
}
func (I *List)Insert(d int, p int) bool {
	if p<0||p>=MAXSIZE||I.length>=MAXSIZE {
		return false
	}
	if p<I.length {
		for k := I.length-1; k>=p; k-- {
			I.Element[k+1]=I.Element[k]
		}
		I.Element[p]=d
		I.length++
		return true
	}else {
		I.Element[I.length]=d
		I.length++
		return true
	}
}

func (I *List)Delete(p int) bool {
	if p<0||p>I.length||p>=MAXSIZE {
		return false
	}
	for ; p<I.length-1; p++ {
		I.Element[p]=I.Element[p+1]
	}
	I.Element[I.length-1]=0
	I.length--
	return true
}