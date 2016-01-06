package binary_tree
import "fmt"

type DataType int
type Node AVLTreeNode
type AVLTree *AVLTreeNode
type AVLTreeNode struct  {
	key DataType
	high int
	left *AVLTreeNode
	right *AVLTreeNode
}

func highTree(p AVLTree) int {
	if p==nil{
		return -1
	}else {
		return p.high
	}
}
func max(a, b int) int {
	if a>b{
		return a
	}else {
		return b
	}
}

//LL
func left_left_rotation(k AVLTree) AVLTree {
	kl:=k.left
	k.left=kl.right
	kl.right=k
	k.high=max(highTree(k.left),highTree(k.right))+1
	kl.high=max(highTree(kl.left),k.high)+1
	return kl
}

//RR
func right_right_rotation(k AVLTree) AVLTree {
	kr:=k.right
	k.right=kr.left
	kr.left=k
	k.high=max(highTree(k.left),highTree(k.right))+1
	kr.high=max(k.high,highTree(kr.right))+1
	return kr
}

func left_right_rotation(k AVLTree) AVLTree {
	k.left=right_right_rotation(k.left)
	return left_left_rotation(k)
}

func right_left_rotation(k AVLTree) AVLTree {
	k.right=left_left_rotation(k.right)
	return right_right_rotation(k)
}

func avl_insert(avl AVLTree, key DataType) AVLTree {
	if avl==nil{
		avl=new(AVLTreeNode)
		if avl==nil{
			fmt.Println("avl tree create error!")
			return nil
		}else {
			avl.key=key
			avl.high=0
			avl.left=nil
			avl.right=nil
		}
	}else if key<avl.key{
		avl.left=avl_insert(avl.left,key)
		if highTree(avl.left)-highTree(avl.right)==2{
			if key<avl.left.key{
				avl=left_left_rotation(avl)
			}else{
				avl=left_right_rotation(avl)
			}
		}
	}else if key>avl.key{
		avl.right=avl_insert(avl.right,key)
		if(highTree(avl.right)-highTree(avl.left))==2{
			if key < avl.right.key { // RL
				avl = right_left_rotation(avl)
			} else {
				fmt.Println("right right", key)
				avl = right_right_rotation(avl)
			}
		}
	}else if key==avl.key{
		fmt.Println("the key",key,"has existed")
	}
	avl.high=max(highTree(avl.left),highTree(avl.right))+1
	return avl
}

func preorder(avl AVLTree) {
	if avl != nil {
		fmt.Print(avl.key, "\t")
		preorder(avl.left)
		preorder(avl.right)
	}
}

func midorder(avl AVLTree) {
	if avl != nil {
		midorder(avl.left)
		fmt.Print(avl.key, "\t")
		midorder(avl.right)
	}
}

/*display avl tree*/
func display(avl AVLTree) {

}