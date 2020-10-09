package main

import "fmt"

func main() {
	cq := CircularQueue{storage:[7]int{}, rear:0, front:0, }
	cq.size = len(cq.storage)

	cq.Enqueue(10)
	cq.Enqueue(20)
	cq.Enqueue(30)
	cq.Enqueue(40)

	cq.Enqueue(100)
	cq.Enqueue(120)

	fmt.Println(cq.Dequeue())
	fmt.Println(cq.Dequeue())
	fmt.Println(cq.Dequeue())
	fmt.Println(cq.Dequeue())
	fmt.Println(cq.Dequeue())
	fmt.Println(cq.Dequeue())
	fmt.Println(cq.Dequeue())
	fmt.Println(cq.Dequeue())
	fmt.Println(cq.Dequeue())
	fmt.Println(cq.Dequeue())
	fmt.Println(cq.Dequeue())
	fmt.Println(cq.Dequeue())
	fmt.Println(cq.Dequeue())
	cq.Enqueue(140)
}

type CircularQueue struct {
	storage [7]int
	rear int
	front int
	size int
}

func (c *CircularQueue) Enqueue(n int){
	if (c.rear+1)%c.size == c.front{
		fmt.Println("Queue is full")
		return
	}else{
		c.rear = (c.rear+1)%c.size
		c.storage[c.rear] = n
	}
}

func (c *CircularQueue) Dequeue()int{
	if (c.front+1)%c.size == c.rear{
		fmt.Println("Queue is empty")
		return -1
	}else{
		c.front = (c.front+1)%c.size
		temp := c.storage[c.front]
		return temp
	}
}





