package linked_list

import "errors"


func (node *Nodestruct) insertLast(data int)error{
	if node.Head == nil{
		return errors.New("Invalid request")
	}
	
	last := node.Head
	for last.Next != nil{
		last = last.Next 
	}
	last.Next = &Node{Data:data, Next:nil}
	return nil
}
