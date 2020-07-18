package main

import "fmt"

func main() {
	q := QueueArr{}
	fmt.Println(q.Pop())
}

type QueueArr struct {
	Storage []int
}

func (stk *QueueArr) Push(data int){
	stk.Storage = append(stk.Storage, data)
}

func (stk *QueueArr) Pop()int{
	first := stk.Storage[0]
	stk.Storage = stk.Storage[1:len(stk.Storage)]
	return first
}

func (stk *QueueArr) Peek()int{
	first := stk.Storage[len(stk.Storage)]
	return first
}






