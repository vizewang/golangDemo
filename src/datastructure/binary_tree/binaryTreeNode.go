package binary_tree
import (
	"math"
)

type BinTreeNode struct {
	data   interface{}
	parent *BinTreeNode
	lChild *BinTreeNode
	rChild *BinTreeNode
	height int
	size   int
}

func NewBinTreeNode(e interface{}) *BinTreeNode {
	return &BinTreeNode{data:e, size:1}
}
func (this *BinTreeNode)GetData() interface{} {
	if this==nil {
		return nil
	}
	return this.data
}
func (this *BinTreeNode)SetData(e interface{}) {
	this.data=e
}

func (this *BinTreeNode)HasParent() bool {
	return this.parent!=nil
}
func (this *BinTreeNode)GetParent() *BinTreeNode {
	if !this.HasParent() {
		return nil
	}
	return this.parent
}

func (this *BinTreeNode) setParent(p *BinTreeNode) {
	this.parent=p
}

func (this *BinTreeNode)CutOffParent() {
	if (!this.HasParent()) {
		return
	}
	if this.IsLChild() {
		this.parent.lChild=nil
	}else {
		this.parent.rChild=nil
	}
}
func (this *BinTreeNode)HasLChild() bool {
	return this.lChild!=nil
}
func (this *BinTreeNode)GetLChild()*BinTreeNode {
	if !this.HasLChild(){
		return nil
	}
	return this.lChild
}

func (this *BinTreeNode)SetLChild(lc *BinTreeNode) *BinTreeNode {
	oldLC:=this.lChild
	if this.HasLChild(){
		this.lChild.CutOffParent()
	}
	if lc!=nil{
		lc.CutOffParent()
		this.lChild=lc
		lc.setParent(this)
		this.SetHeight()
		this.SetSize()
	}
	return oldLC
}

func (this *BinTreeNode) HasRChild() bool {
	return this.rChild != nil
}


func (this *BinTreeNode) GetRChild() *BinTreeNode {
	if !this.HasRChild() {
		return nil
	}
	return this.rChild
}

func (this *BinTreeNode) SetRChild(rc *BinTreeNode) *BinTreeNode {
	oldRC := this.rChild
	if this.HasRChild() {
		this.rChild.CutOffParent()
	}
	if rc != nil {
		rc.CutOffParent()
		this.rChild = rc
		rc.setParent(this)
		this.SetHeight()
		this.SetSize()
	}
	return oldRC
}

func (this *BinTreeNode)GetHeight() int {
	return this.height
}
func (this *BinTreeNode)SetHeight() {
	newH:=0
	if this.HasLChild(){
		newH=int(math.Max(float64(newH),float64(1+this.GetLChild().GetHeight())))
	}
	if this.HasRChild(){
		newH=int(math.Max(float64(newH),float64(1+this.GetRChild().GetHeight())))
	}
	if newH==this.height{
		return
	}
	this.height=newH
	if this.HasParent(){
		this.GetParent().SetHeight()
	}
}

func (this *BinTreeNode)GetSize() int {
	return this.size
}

func (this *BinTreeNode) SetSize() {
	this.size = 1 //初始化为1,结点本身
	if this.HasLChild() {
		this.size += this.GetLChild().GetSize() //加上左子树规模
	}
	if this.HasRChild() {
		this.size += this.GetRChild().GetSize() //加上右子树规模
	}

	if this.HasParent() {
		this.parent.SetSize() //递归更新祖先的规模
	}

}

func (this *BinTreeNode)IsLChild() bool {
	return this.HasParent()&&this==this.parent.lChild
}
func (this *BinTreeNode)IsRChild() bool {
	return this.HasParent()&&this==this.parent.rChild
}