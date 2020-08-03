package main

import "fmt"

func main() {
    q := Queue{}
	treeNode := TreeNode{q:&q}
	treeNode.InsertBinaryTree(10, -1)
	treeNode.InsertBinaryTree(20, 30)
	treeNode.InsertBinaryTree(25, 27)
	treeNode.InsertBinaryTree(35, 37)
	treeNode.InsertBinaryTree(92, 100)
	treeNode.InsertBinaryTree(95, 101)

	//treeNode.PreOrder(treeNode.Root)

	//treeNode.PostOrder(treeNode.Root)

	//treeNode.InOrder(treeNode.Root)

	//treeNode.LevelOrder(treeNode.Root)

	//fmt.Print(treeNode.CountNodes(treeNode.Root))
	//fmt.Println(treeNode.CountFullNodes(treeNode.Root))
	fmt.Println(treeNode.CountLeafNodes(treeNode.Root))
}

type Tree struct {
	Data int
	LeftChild  *Tree
	RightChild *Tree
}

type TreeNode struct {
	Root *Tree
	q *Queue
	counter int
}

func (tree *TreeNode) InsertBinaryTree(leftData int, rightData int){
	if tree.Root == nil{
		tree.Root = &Tree{Data:leftData, LeftChild:nil, RightChild:nil}
		tree.q.Enqueue(tree.Root)
		return
	}
	if !tree.q.IsEmpty(){
		data :=tree.q.Dequeue()
		if data.LeftChild == nil && leftData != -1{
			data.LeftChild = &Tree{Data:leftData, LeftChild:nil, RightChild:nil}
			tree.q.Enqueue(data.LeftChild)
		}
		if data.RightChild == nil && rightData != -1{
			data.RightChild = &Tree{Data:rightData, LeftChild:nil, RightChild:nil}
			tree.q.Enqueue(data.RightChild)
		}
	}
}

func (tree *TreeNode) PreOrder(root *Tree){
	if root == nil{
		return
	}
	fmt.Println(root.Data)
	tree.PreOrder(root.LeftChild)
	tree.PreOrder(root.RightChild)
}

func (tree *TreeNode) PostOrder(root *Tree){
	if root == nil{
		return
	}
	tree.PostOrder(root.LeftChild)
	tree.PostOrder(root.RightChild)
	fmt.Println(root.Data)
}

func (tree *TreeNode) InOrder(root *Tree){
	if root == nil{
		return
	}
	tree.InOrder(root.LeftChild)
	fmt.Println(root.Data)
	tree.InOrder(root.RightChild)
}

func (tree *TreeNode) LevelOrder(root *Tree){
	if root == nil{
		return
	}
	q := Queue{}
	q.Enqueue(root)
	for !q.IsEmpty(){
		data := q.Dequeue()
		fmt.Println(data.Data)
		if data.LeftChild != nil{
			q.Enqueue(data.LeftChild)
		}
		if data.RightChild != nil{
			q.Enqueue(data.RightChild)
		}
	}
}

func (tree *TreeNode) CountNodes(root *Tree)int{
	if root == nil{
		return 0
	}
	x := tree.CountNodes(root.LeftChild)
	y := tree.CountNodes(root.RightChild)
	tree.counter = 1+x+y
	return tree.counter
}

func (tree *TreeNode) CountFullNodes(root *Tree)int{
	if root == nil{
		return 0
	}
	if root.LeftChild != nil && root.RightChild != nil{
		x := tree.CountFullNodes(root.LeftChild)
		y := tree.CountFullNodes(root.RightChild)
		return 1+x+y
	}else{
		return 0
	}

}


func (tree *TreeNode) CountLeafNodes(root *Tree)int{
	if root == nil{
		return 0
	}
	if root.LeftChild != nil && root.RightChild != nil{
		x := tree.CountLeafNodes(root.LeftChild)
		y := tree.CountLeafNodes(root.RightChild)
		return x+y
	}else{
		return 1
	}

}




type Queue struct {
	Storage []*Tree
}

func (q *Queue) Enqueue(data *Tree){
	q.Storage = append(q.Storage, data)
}

func (q *Queue) Dequeue()*Tree{
	if q.Storage == nil{
		return nil
	}
	temp := q.Storage[0]
	q.Storage = q.Storage[1:len(q.Storage)]
	return temp
}

func (q *Queue) IsEmpty()bool{
	return len(q.Storage) <= 0
}


