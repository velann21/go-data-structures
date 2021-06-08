package stack

type Node struct {
	data int
	next *Node
}

func NewNode(data int, next *Node)*Node{
	return &Node{data: data, next: next}
}

type Stack struct {
	Size int
	top *Node
}

func NewStackLL()*Stack{
	return &Stack{}
}

func (ll *Stack) Push(data int){
	if ll.top == nil{
		ll.top = NewNode(data, nil)
		ll.Size++
		return
	}
	newNode := NewNode(data, ll.top)
	ll.top = newNode
	ll.Size++
}

func (ll *Stack) Pop()(int, error){
	if ll.top == nil{
		return 0, NoData
	}
	data := ll.top.data
	if ll.top.next != nil{
		ll.top = ll.top.next
	}
	ll.Size--
	return data, nil
}

func (ll *Stack) Peek(pos int)(int, error){
	if ll.top == nil{
		return 0, NoData
	}
	if pos <=0 || pos > ll.Size{
		return 0, InvalidIndex
	}
	if pos == 1{
		return ll.top.data, nil
	}
	last := ll.top
	currentPos := 1
	for last != nil {
		if pos == currentPos{
			return last.data, nil
		}
		last = last.next
		currentPos++
	}
	return 0, nil
}

func (ll *Stack) QueueUsingStack(auxilaryStack *Stack, data int)error{
	if ll.Size != 0 && auxilaryStack == nil{
		return InvalidData
	}
	if data < 0{
		return InvalidData
	}
	if ll.Size == 0{
		ll.Push(data)
	}else {
		for ll.Size>0{
			data, err := ll.Pop()
			if err != nil{
				return InvalidData
			}
			auxilaryStack.Push(data)
		}
		ll.Push(data)
		for auxilaryStack.Size>0{
			data, err := auxilaryStack.Pop()
			if err != nil{
				return InvalidData
			}
			ll.Push(data)
		}

	}
	return nil
}
