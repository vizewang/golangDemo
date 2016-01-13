package binary_tree
import (
	"testing"
	"fmt"
)


func TestTree(t *testing.T) {
	a:=NewBinTreeNode(1)
	tree1:=NewBinaryTree(a)
	a.SetLChild(NewBinTreeNode(2))
	a.SetRChild(NewBinTreeNode(5))
	a.GetLChild().SetRChild(NewBinTreeNode(3))
	a.GetLChild().GetRChild().SetLChild(NewBinTreeNode(4))
	a.GetRChild().SetLChild(NewBinTreeNode(6))
	a.GetRChild().SetRChild(NewBinTreeNode(7))
	a.GetRChild().GetLChild().SetRChild(NewBinTreeNode(9))
	a.GetRChild().GetRChild().SetRChild(NewBinTreeNode(8))

	node2 := a.GetLChild()
	node9 := a.GetRChild().GetLChild().GetRChild()

	if node2.IsLeaf(){
		t.Error("node2 is not leaf")
	}
	if !node9.IsLeaf(){
		t.Error("node9 is a leaf")
	}

	if tree1.GetSize()!=9{
		t.Error("size is 9")
	}

	I :=tree1.InOrder()
	for e:=I.Front();e!=nil;e=e.Next(){
		obj,_:=e.Value.(*BinTreeNode)
		t.Logf("data:%v\n",*obj)
	}
	result := tree1.Find(6)
	fmt.Printf("结点6的父节点数据：%v\t结点6的右孩子节点数据: %v\n", result.GetParent().GetData(), result.GetRChild().GetData())
}
