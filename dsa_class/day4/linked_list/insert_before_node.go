package linked_list

import (
	"errors"
	"fmt"
)

func main() {
	node := Nodestruct1{}
	node.Insert(10)
	node.Insert(20)
	node.Insert(30)
	node.Insert(40)
	node.Insert(50)
	node.Insert(60)
	err := node.addBefore(5, 10)
	if err != nil{
		fmt.Println(err)
	}

}

type Node1 struct {
	Data int
	Next *Node1
}

type Nodestruct1 struct {
	Head *Node1
}

func (node *Nodestruct1) Insert(data int){
	if node.Head == nil{
		node.Head = &Node1{Data:data, Next:nil}
		return
	}

	last := node.Head
	for last.Next != nil{
		last = last.Next
	}
	last.Next = &Node1{Data:data, Next:nil}
}

func (node *Nodestruct1) addBefore(data1 int, data2 int)error{
	if node.Head == nil{
		return errors.New("Invalid Request")
	}
	if node.Head.Data == data2{
		temp := node.Head
		node.Head = &Node1{Data:data1,Next:temp}
		return nil
	}
	last := node.Head
	previous := &Node1{}
	for last != nil{
		if last.Data == data2{
			temp := previous.Next
			previous.Next = &Node1{Data:data1, Next:temp}
			break
		}
		previous = last
		last = last.Next
	}
	return nil
}


