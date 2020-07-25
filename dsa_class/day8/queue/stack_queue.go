package queue

import "fmt"

func main() {
	node := queueNodeStruct{}
	node1 := queueNodeStruct{}
	node.FormStackInQueue(10, &node1)
	node.FormStackInQueue(20, &node1)
	node.FormStackInQueue(30, &node1)
	node.FormStackInQueue(40, &node1)

	fmt.Println(node.Dequeue().Data)
	fmt.Println(node.Dequeue().Data)
}

type queueNode struct {
	Data int
	Next *queueNode
}

type queueNodeStruct struct {
	Head *queueNode
	Tail *queueNode
}

func (q *queueNodeStruct) Enqueue(data int){
	if q.Head == nil{
		q.Head = &queueNode{Data:data, Next:nil}
		q.Tail = q.Head
		return
	}
	newNode := &queueNode{Data:data, Next:nil}
	q.Tail.Next = newNode
	q.Tail = newNode
	return
}

func (q *queueNodeStruct) Dequeue()*queueNode{
	if q.Head == nil{
		return nil
	}
	temp :=  q.Head
	if q.Head.Next != nil{
		q.Head = q.Head.Next
	}else{
		q.Head = nil
	}
	return temp
}

func (q *queueNodeStruct) FormStackInQueue(data int, temp *queueNodeStruct){
	if q.Head == nil{
		q.Enqueue(data)
		return
	}
	qData := q.Dequeue()
	for qData != nil{
		temp.Enqueue(qData.Data)
		qData = q.Dequeue()
	}
	q.Enqueue(data)

	temdata := temp.Dequeue()
	for temdata != nil{
		q.Enqueue(temdata.Data)
		temdata = temp.Dequeue()
	}
}
