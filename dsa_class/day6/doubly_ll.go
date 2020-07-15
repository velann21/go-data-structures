package main

func main() {
	node := NodeStruct{}
	node.Insert(10)
	node.Insert(20)
	node.Insert(30)
	node.Insert(40)
	node.Insert(50)

	node.AddToHead(05)
	node.AddBefore(20, 15)
	node.AddAfter(20, 25)

	node.RemoveHead()
	node.RemoveEnd()
}

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type NodeStruct struct {
	Head *Node
}

func (node *NodeStruct) Insert(data int){

	if node.Head == nil{
		node.Head = &Node{Data:data, Next:nil, Prev:nil}
		return
	}

	last := node.Head

	for last.Next != nil{
		last = last.Next
	}

	last.Next = &Node{Data:data, Next:nil, Prev:last}
	return
}

func (node *NodeStruct) AddToHead(data int){
	if node.Head == nil{
		node.Head = &Node{Data:data, Next:nil, Prev:nil}
		return
	}
	temp := node.Head
	node.Head = &Node{Data:data, Next:temp, Prev:nil}
	temp.Prev = node.Head
	return
}

func (node *NodeStruct) AddBefore(data1 int, data2 int){
	if node.Head == nil{
		return
	}

	if node.Head.Data == data1{
		node.AddToHead(data2)
	}

	last := node.Head
	for last != nil{
		last = last.Next
		if last.Data == data1{
			tempprev := last.Prev
			newNode := &Node{Data:data2, Next:last, Prev:tempprev}
			tempprev.Next = newNode
			last.Prev = newNode
			return
		}

	}
}


func (node *NodeStruct) AddAfter(data1 int, data2 int){
	if node.Head == nil{
		return
	}

	last := node.Head
	for last != nil{
		last = last.Next
		if last.Data == data1{
			tempprev := last.Next
			newNode := &Node{Data:data2, Next:tempprev, Prev:last}
			tempprev.Prev = newNode
			last.Next = newNode
			break
		}

	}
}

func (node *NodeStruct) RemoveHead(){
	if node.Head == nil{
		return
	}
	temp := node.Head
	if temp.Next != nil{
		node.Head = temp.Next
	}
}

func (node *NodeStruct) RemoveEnd(){
	if node.Head == nil{
		return
	}
	last := node.Head
	for last.Next != nil{
		last = last.Next
	}

	last = last.Prev
	last.Next = nil
}

