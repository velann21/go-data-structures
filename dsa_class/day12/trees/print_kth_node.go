package main

import "fmt"

func main() {

	tree := TreeTraverser{}
	tree.Insert(20)
	tree.Insert(10)
	tree.Insert(30)
	tree.Insert(05)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(35)
	tree.Insert(8)

	tree.PrintKthNode(tree.Root, 3)

}

type TreeNode1 struct {
	Data int
	LeftChild *TreeNode1
	RightChild *TreeNode1
}

type TreeTraverser struct {
	Root *TreeNode1
}

func (tree *TreeTraverser) Insert(data int){
	if tree.Root == nil{
		tree.Root = &TreeNode1{Data:data, LeftChild:nil, RightChild:nil}
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
		parentNode.LeftChild = &TreeNode1{Data:data, LeftChild:nil, RightChild:nil}
	}else if data > parentNode.Data{
		parentNode.RightChild = &TreeNode1{Data:data, LeftChild:nil, RightChild:nil}
	}
}

func (tree *TreeTraverser) PrintKthNode(root *TreeNode1, k int){
	if root == nil{
		return
	}
	if k==1{
		fmt.Println(root.Data)
		return
	}
	tree.PrintKthNode(root.LeftChild, k-1)
	tree.PrintKthNode(root.RightChild, k-1)
}
