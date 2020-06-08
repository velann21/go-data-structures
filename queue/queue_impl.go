package queue

import "fmt"

func main() {

	q := New()
	fmt.Println(q.IsEmpty())
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)
	q.Enqueue(60)
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Front())
	fmt.Println(q.Rear())

	devalue := q.Dequeue()
	fmt.Println(devalue)

}

type Queue struct {
	Items []int
}

func New() *Queue {
	arr := make([]int, 0)
	newQueue := Queue{Items:arr}
	return &newQueue
}

func (q *Queue) Enqueue(value int){
	q.Items = append(q.Items, value)
}

func (q *Queue) Dequeue()int{
	data := q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	return data
}

func (q *Queue) Front() int{
	data := q.Items[0]
	return data
}

func (q *Queue) Rear() int {
	data := q.Items[len(q.Items)-1]
	return data
}

func (q *Queue) IsEmpty()bool{
	return len(q.Items) == 0
}


