package main

import "fmt"

func main() {
	tree := TreeTraverser4{}
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(30)
	tree.Insert(10)
	tree.Insert(18)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(25)
	tree.Insert(35)
	tree.Insert(22)
	tree.Insert(45)
	tree.Insert(44)
	tree.Insert(55)
	tree.PostOrderTraverser(tree.Root)

	fmt.Println(tree.CountHalfLeafNodes(tree.Root))

}

type TreeNode4 struct {
	Data int
	LeftChild *TreeNode4
	RightChild *TreeNode4
}

type TreeTraverser4 struct {
	Root *TreeNode4
}

func (tree *TreeTraverser4) Insert(data int){
	if tree.Root == nil{
		tree.Root = &TreeNode4{Data:data, LeftChild:nil, RightChild:nil}
		return
	}
	temp := tree.Root
	parentNode := temp
	for temp != nil{
		parentNode = temp
		if data < temp.Data{
			temp = temp.LeftChild
		}else if data > temp.Data{
			temp = temp.RightChild
		}
	}

	if data < parentNode.Data{
		parentNode.LeftChild = &TreeNode4{Data:data, LeftChild:nil, RightChild:nil}
	}else if data > parentNode.Data{
		parentNode.RightChild = &TreeNode4{Data:data, LeftChild:nil, RightChild:nil}
	}
}

func (tree *TreeTraverser4) PostOrderTraverser(root *TreeNode4){
	if root == nil{
		return
	}
	tree.PostOrderTraverser(root.LeftChild)
	tree.PostOrderTraverser(root.RightChild)
	fmt.Println(root.Data)

}

func (tree *TreeTraverser4) CountHalfLeafNodes(root *TreeNode4)int{
	if root == nil{
		return 0
	}
	lres := 0
	rres := 0
	res := 0
	if root.LeftChild != nil && root.RightChild == nil{
		lres = lres +1
	}else if root.RightChild != nil && root.LeftChild == nil{
		rres = rres +1
	}
	x := lres + tree.CountHalfLeafNodes(root.LeftChild)
	y := rres + tree.CountHalfLeafNodes(root.RightChild)
	res = x+y
	return res
}