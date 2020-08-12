package main

import "fmt"

func main() {
	tree := TreeTraverser4{}
    tree.Insert(20)
	tree.Insert(10)
	tree.Insert(30)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(35)

	fmt.Println(tree.LCA(tree.Root , 5, 35).Data)
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
	lastParent := temp
	for temp != nil{
		lastParent = temp
		if data < temp.Data{
			temp = temp.LeftChild
		}else if data > temp.Data{
			temp = temp.RightChild
		}
	}
	if data < lastParent.Data{
		lastParent.LeftChild = &TreeNode4{Data:data, LeftChild:nil, RightChild:nil}
	}else if data > lastParent.Data{
		lastParent.RightChild = &TreeNode4{Data:data, LeftChild:nil, RightChild:nil}
	}
}

func (tree *TreeTraverser4) LCA(root *TreeNode4, data1, data2 int)*TreeNode4{
	if root == nil{
		return nil
	}
	if root.Data == data1 || root.Data == data2{
		return root
	}
	left := tree.LCA(root.LeftChild, data1, data2)
	right := tree.LCA(root.RightChild, data1, data2)
	if left != nil && right != nil{
		return root
	}else if left != nil{
		return left
	}else if right != nil{
		return right
	}else {
		return nil
	}
}
