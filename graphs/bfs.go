package main

import (
	q "awesomeProject3/queue"
	"fmt"
)

func main() {

	// Type1
	am := [][]int{
		{0,0,0,0,0,0,0},
		{0,0,1,1,0,0,0},
		{0,1,0,0,1,0,0},
		{0,1,0,0,0,0,0},
		{0,0,1,1,0,1,1},
		{0,0,0,0,1,0,0},
		{0,0,0,0,1,0,0},
	}
	BFS(am, 6, 7)

    //Type2
	node := NodeItems{}
	nA := node.AddNode("A")
	nB := node.AddNode("B")
	nc := node.AddNode("C")
	nD := node.AddNode("D")
	nE := node.AddNode("E")
	nF := node.AddNode("F")
	nG := node.AddNode("G")

	node.AddEdges(nA, nB)
	node.AddEdges(nA, nc)
	node.AddEdges(nB, nc)
	node.AddEdges(nc, nD)
	node.AddEdges(nD, nA)
	node.AddEdges(nD, nE)
	node.AddEdges(nE, nF)
	node.AddEdges(nE, nG)
	node.AddEdges(nF, nG)

	node.PrintOutEdges()

	node.Bfs()


}

func BFS (graph [][]int, startPoint int, verticesSize int){

	fmt.Print(startPoint,",")
	visited := make([]int, verticesSize)
	visited[startPoint] = 1
	queue := q.New()

	queue.Enqueue(startPoint)

	for !queue.IsEmpty() {
		vertice := queue.Dequeue()
		for v:=1; v<verticesSize; v++{
			if graph[vertice][v]==1 && visited[v] == 0 {
				fmt.Print(v, ",")
				visited[v] = 1
				queue.Enqueue(v)
			}
		}
	}





}



type Node1 struct {
	Item string
}

type NodeItems struct {
	Items []*Node1
	OutDegree map[*Node1][]*Node1
}

func (nodeItems *NodeItems) AddNode(ele string) *Node1 {
	element := Node1{Item:ele}
	nodeItems.Items = append(nodeItems.Items, &element)
	return &element
}


func (nodeItems *NodeItems) AddEdges(n1 *Node1, n2 *Node1){
	if nodeItems.OutDegree == nil{
		nodeItems.OutDegree = make(map[*Node1][]*Node1)
	}

	nodeItems.OutDegree[n1] = append(nodeItems.OutDegree[n1], n2)
	nodeItems.OutDegree[n2] = append(nodeItems.OutDegree[n2], n1)
}


func (nodeItems *NodeItems) PrintOutEdges(){
	for _, v := range nodeItems.Items{
		fmt.Print(v.Item, "-> ")
		outV := nodeItems.OutDegree[v]
		for _, oV := range  outV {
			fmt.Print(oV.Item, " ")
		}
		fmt.Println()
	}
}

func (nodeItems *NodeItems) Bfs(){
	q := Queue{}
	q.enqueue(nodeItems.Items[2])
	visited := make(map[*Node1]bool)
	visited[nodeItems.Items[2]] = true


	fmt.Print(nodeItems.Items[2], ",")
	for !q.IsEmpty(){
		vertices := q.dequeue()
		for _, v := range nodeItems.OutDegree[vertices]{
			if !visited[v]{
				fmt.Print(v, ",")
				q.enqueue(v)
				visited[v] = true
			}
		}
	}
}


type Queue struct {
	Queue []*Node1
}

func (q *Queue) enqueue(item *Node1){
	q.Queue = append(q.Queue, item)
}

func (q *Queue) dequeue() *Node1{
	one := q.Queue[0]
	q.Queue = q.Queue[1:len(q.Queue)]
	return one
}

func (q *Queue) peek() *Node1 {
	one := q.Queue[0]
	return one
}

func (q *Queue) IsEmpty() bool {
	return len(q.Queue) == 0
}

func (q *Queue) Front() *Node1{
	data := q.Queue[0]
	return data
}

func (q *Queue) Rear() *Node1 {
	data := q.Queue[len(q.Queue)-1]
	return data
}



