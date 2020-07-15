package main


func main(){


}

type NodePtr struct {
	Data int
	Next *NodePtr
}

type NodePtrStruct struct {
	Head *NodePtr
}

func (node *NodePtrStruct) Insert(data int){
	if node.Head == nil{
		node.Head = &NodePtr{Data:data, Next:nil}
		return
	}

	last := node.Head

	for last.Next != nil{
		last = last.Next
	}

	last.Next = &NodePtr{Data:data, Next:nil}
	return
}

func (node *NodePtrStruct) addTwoNumbers(node2 *NodePtrStruct) *NodePtr {
	if node == nil || node2 == nil{
		return nil
	}
	n1 := node.Head
	n2 := node2.Head
	sum := 0
	for n1 != nil && n2 != nil{
		sum = n1.Data + n2.Data
		first := sum / 10;
		last := sum % 10;
		n1.Data = first
		if n1.Next != nil{
			n1.Next.Data += last
		}
		n1 = n1.Next
		n2 = n2.Next
	}

	return n1

}
