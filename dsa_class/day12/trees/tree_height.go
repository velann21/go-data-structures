package main

import "fmt"

func main() {
	tree := TreeNodeTraverser{}
	tree.Insert(20)
	tree.Insert(10)
	tree.Insert(30)
	tree.Insert(05)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(35)

	fmt.Println(tree.FindHeight(tree.Root))
}

type TreeNode struct {
	Data int
	LeftChild *TreeNode
	RightChild *TreeNode
}

type TreeNodeTraverser struct {
	Root *TreeNode
}

func (tree *TreeNodeTraverser) Insert(data int){
	if tree.Root == nil{
		tree.Root = &TreeNode{Data:data, LeftChild:nil, RightChild:nil}
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
		parentNode.LeftChild = &TreeNode{Data:data, LeftChild:nil, RightChild:nil}
	}else if data > parentNode.Data{
		parentNode.RightChild = &TreeNode{Data:data, LeftChild:nil, RightChild:nil}
	}
}

func (tree *TreeNodeTraverser) FindHeight(root *TreeNode)int64{
	if root == nil{
		return 0
	}
	x := 1+tree.FindHeight(root.LeftChild)
	y := 1+tree.FindHeight(root.RightChild)
	return Max(x, y)
}

func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}


