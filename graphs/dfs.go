package main

import "fmt"

func main() {


	node := NodeItems2{}
	nA := node.AddNode("A")
	nB := node.AddNode("B")
	nC := node.AddNode("C")
	nD := node.AddNode("D")
	nE := node.AddNode("E")
	nF := node.AddNode("F")

	node.AddEdges(nA, nB)
	node.AddEdges(nA, nC)
	node.AddEdges(nB, nD)
	node.AddEdges(nC, nD)
	node.AddEdges(nD, nE)
	node.AddEdges(nD, nF)

	node.PrintAllOutEdges()
	node.Bfs(nF)

	node.Dfs(nA)


}

type Node2 struct {
	Item string
}

type NodeItems2 struct {
	Items []*Node2
	OutDegree map[*Node2][]*Node2
	visited map[*Node2]bool
}

func (node *NodeItems2) AddNode(data string) *Node2 {
	nodeIt := Node2{Item:data}
	node.Items = append(node.Items, &nodeIt)
	return &nodeIt
}

func (node *NodeItems2) AddEdges(n1 *Node2, n2 *Node2){

	if node.OutDegree == nil{
		node.OutDegree = make(map[*Node2][]*Node2)
	}
	node.OutDegree[n1] = append(node.OutDegree[n1], n2)
	node.OutDegree[n2] = append(node.OutDegree[n2], n1)
}

func (node *NodeItems2) PrintAllOutEdges(){
	for _,v := range node.Items{
		fmt.Print(v.Item, "-> ")
		for _, v1 := range node.OutDegree[v]{
			fmt.Print(v1.Item, " ")
		}
		fmt.Println()
	}
}

func (node *NodeItems2) Bfs(starting *Node2){
	q := Queue1{}
	visitedItem := make(map[*Node2]bool)
	q.enqueue(starting)
	visitedItem[starting] = true
	fmt.Print(starting ,",")
	for !q.IsEmpty(){
		vertice := q.dequeue()
		for _, v := range node.OutDegree[vertice]{
			if !visitedItem[v]{
				fmt.Print(v ,",")
				q.enqueue(v)
				visitedItem[v] = true
			}
		}
	}
	fmt.Println()
}


func (node *NodeItems2) Dfs(starting *Node2){
	if node.visited == nil{
		node.visited = make(map[*Node2]bool)
	}
	if !node.visited[starting]{
		fmt.Print(starting)
		node.visited[starting] = true
		for _, v := range node.OutDegree[starting]{
            if !node.visited[v]{
				//fmt.Print(starting)
            	node.Dfs(v)
			}
		}
	}



}


type Queue1 struct {
	QueueItem []*Node2
}

func (q *Queue1) enqueue(data *Node2){
	q.QueueItem = append(q.QueueItem, data)
}

func (q *Queue1) dequeue() *Node2 {
	popFirst := q.QueueItem[0]
	q.QueueItem = q.QueueItem[1:len(q.QueueItem)]
	return popFirst
}

func (q *Queue1) peek()*Node2{
	peekFirst := q.QueueItem[0]
	return peekFirst
}

func (q *Queue1) IsEmpty() bool {
	return len(q.QueueItem) == 0
}

func (q *Queue1) front() *Node2 {
	return  q.QueueItem[0]
}

func (q *Queue1) rear() *Node2 {
	return q.QueueItem[len(q.QueueItem)-1]
}


type Stack struct {
	StackItem []*Node2
}

func (stk *Stack) push(node *Node2){
	stk.StackItem = append(stk.StackItem, node)
}

func (stk *Stack) pop() *Node2 {
	item := stk.StackItem[len(stk.StackItem)-1]
	stk.StackItem = stk.StackItem[0:len(stk.StackItem)-1]
	return item
}

func (stk *Stack) peek() *Node2 {
	popedItem := stk.StackItem
	item := popedItem[len(popedItem)-1]
	return item
}




