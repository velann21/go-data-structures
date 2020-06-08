package main

import "fmt"

func main(){
	node := NodeItem4{}
	nA := node.AddNode("A")
	nB := node.AddNode("B")
	nC := node.AddNode("C")
	nD := node.AddNode("D")
	nE := node.AddNode("E")

	node.AddEdge(nA, nB)
	node.AddEdge(nB, nA)

	node.AddEdge(nA, nC)
	node.AddEdge(nC, nA)

	node.AddEdge(nB, nD)
	node.AddEdge(nD, nB)

	node.AddEdge(nB, nE)
	node.AddEdge(nE, nB)

	node.AddEdge(nB, nC)
	node.AddEdge(nC, nB)

	//node.AddEdge(nE, nD)
	//node.AddEdge(nD, nE)

	node.PrintAllEdges()
	node.Bfs(nA)

	visitedMap := make(map[Node4]int)
	visitedMap[nA] = -1
	visitedMap[nB] = -1
	visitedMap[nC] = -1
	visitedMap[nD] = -1
	visitedMap[nE] = -1
	node.DetectCycle(nA, visitedMap)



}

type Node4 struct {
	Element string

}

type NodeItem4 struct {
	Items []Node4
	OutDegrees map[Node4][]Node4
}

func (node *NodeItem4) AddNode(value string) Node4{
    nodeIt := Node4{Element:value}
	node.Items = append(node.Items, nodeIt)
	return nodeIt
}

func (node *NodeItem4) AddEdge(n1 Node4, n2 Node4){
	if node.OutDegrees == nil{
		node.OutDegrees = make(map[Node4][]Node4)
	}
	node.OutDegrees[n1] = append(node.OutDegrees[n1], n2)
}

func (node *NodeItem4) PrintAllEdges(){
	for _, v := range node.Items{
		fmt.Print(v, "-> ")
		for _, v1 := range node.OutDegrees[v]{
			fmt.Print(v1)
		}
		fmt.Println()
	}
}

func (node *NodeItem4) DetectCycle(starting Node4,visitedMap map[Node4]int){

	q := Queue4{}
	q.enqueue(starting)
	visitedMap[starting] = 0

	for !q.isEmpty(){
		vertices := q.dequeue()
		visitedMap[vertices] = 1
		for _, v := range node.OutDegrees[vertices]{
			if visitedMap[v] == 0{
				fmt.Println("Cycle detected")
				fmt.Println(vertices)
				fmt.Println(node.OutDegrees[vertices])
				return
			}
            if visitedMap[v] != 1{
            	q.enqueue(v)
				visitedMap[v] = 0
			}
		}
	}
}

func (node *NodeItem4) Bfs(starting Node4){
	visitedMap := map[Node4]bool{}
	q := Queue4{}
	q.enqueue(starting)
	visitedMap[starting] = true
	fmt.Println(starting, ",")
	for !q.isEmpty(){
		vertices := q.dequeue()
		for _, v := range node.OutDegrees[vertices]{
			if !visitedMap[v]{
				fmt.Println(v, ",")
				visitedMap[v] = true
				q.enqueue(v)
			}
		}
	}

}



type Queue4 struct {
	Queue []Node4
}

func (q *Queue4) enqueue(data Node4){
	q.Queue = append(q.Queue, data)
}

func (q *Queue4) dequeue()Node4{
	popFirst := q.Queue[0]
	q.Queue = q.Queue[1:len(q.Queue)]
	return popFirst
}

func (q *Queue4) isEmpty() bool{
	return len(q.Queue) == 0
}

