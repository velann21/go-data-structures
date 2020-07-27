package linked_list

import "fmt"

func main() {
	node := NodeStruct19{}
	node.Insert(10)
	node.Insert(20)
	node.Insert(30)
	node.Insert(40)
	node.Insert(50)
	node.Insert(60)
	node.Insert(70)
	node.Insert(80)
	node.Insert(90)
	node.Insert(100)

	node.FindKthElement()
}

type Node19 struct {
	Data int
	Next *Node19
}

type NodeStruct19 struct {
	Root *Node19
}

func (node *NodeStruct19) Insert(data int){
	if node.Root == nil{
		node.Root = &Node19{Data:data, Next:nil}
		return
	}
	last := node.Root
	for last.Next != nil{
		last = last.Next
	}
	last.Next = &Node19{Data:data, Next:nil}
}

func (node *NodeStruct19) FindKthElement(){
	if node.Root == nil{
		return
	}

	last := node.Root
	fast := node.Root
	fast = movektime(fast, 5)
	for fast != nil{
		fast = fast.Next
		last = last.Next
	}
	fmt.Print(last.Data)
}

func movektime(root *Node19, moves int) *Node19{
	counter := 0
	for root != nil && counter < moves{
		root = root.Next
		counter ++
	}
	return root
}