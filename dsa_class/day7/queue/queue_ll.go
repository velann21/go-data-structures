package main

import "fmt"

func main(){
	q := QueueNodeStruct{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)

	fmt.Println(q.Dequeue().Data)
	fmt.Println(q.Dequeue().Data)

}

type QueueNode struct {
	Data int
	Next *QueueNode
}

type QueueNodeStruct struct {
	Head *QueueNode
}

func (q *QueueNodeStruct) Enqueue(data int){
	if q.Head == nil{
		q.Head = &QueueNode{Data:data, Next:nil}
		return
	}

	last := q.Head
	for last.Next != nil{
		last = last.Next
	}

	last.Next = &QueueNode{Data:data, Next:nil}
}

func (q *QueueNodeStruct) Dequeue()*QueueNode{
	if q.Head == nil{
		return nil
	}
	temp := q.Head
	q.Head = q.Head.Next
	return temp
}