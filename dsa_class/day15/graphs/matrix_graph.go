package main

import "fmt"

func main() {

	gu := GraphUtil{}
	gu.AddNode("A")
	gu.AddNode("B")
	gu.AddNode("C")
	gu.AddNode("D")
	gu.AddNode("E")

	gu.AddEgdes("A", "B")
	gu.AddEgdes("A", "E")
	gu.AddEgdes("B", "C")
	gu.AddEgdes("C", "D")
	gu.AddEgdes("D", "E")
	fmt.Println(gu.matrix)
	gu.PrintGraph()
	gu.BFS("A")
}

type GraphNode struct {
	Position int
	Label string
}

type GraphUtil struct {
	matrix [5][5]int
	VertexTracker map[string]*GraphNode
	InverseVertexTracker map[int]*GraphNode
	PositionTracker int
}

func (g *GraphUtil) AddNode(data string){
	if data == ""{
		return
	}
	if g.VertexTracker == nil{
		g.VertexTracker = make(map[string]*GraphNode, 0)
		g.InverseVertexTracker = make(map[int]*GraphNode, 0)
	}
	gn := GraphNode{Position:g.PositionTracker, Label:data}
	g.InverseVertexTracker[g.PositionTracker] = &gn
	g.PositionTracker = g.PositionTracker+1
	g.VertexTracker[data] = &gn
}

func (g *GraphUtil) AddEgdes(node1 string, node2 string){
	if node1 == "" || node2 == ""{
		return
	}
	fromNodeVertex := g.VertexTracker[node1]
	toNodeVertex := g.VertexTracker[node2]
	if fromNodeVertex == nil || toNodeVertex == nil{
		return
	}
	g.matrix[fromNodeVertex.Position][toNodeVertex.Position] = 1
}

func (g *GraphUtil) PrintGraph(){
	for i, v := range g.matrix{
		Vertex := g.InverseVertexTracker[i]
		fmt.Print(Vertex.Label, "----->")
		for j, arrayLooper := range v{
			if arrayLooper == 1{
				fmt.Print(g.InverseVertexTracker[j].Label, ",")
			}
		}
		fmt.Println()
	}
}

func (g *GraphUtil) BFS(start string){
	if g.VertexTracker == nil{
		return
	}
	if start == ""{
		return
	}
	visitedMap := map[string]bool{}
	startingNode := g.VertexTracker[start]
	sn := QueueUtil{}
	visitedMap[startingNode.Label] = true
	sn.Enqueue(startingNode)
	fmt.Println(startingNode.Label)
	for !sn.IsEmpty(){
		element := sn.Dequeue()
		elementPosition := g.VertexTracker[element.Label].Position
		for i, edges := range g.matrix{
			if i != elementPosition{
				continue
			}
			for j, edge := range edges{
				edgeInStr := g.InverseVertexTracker[j]
				if edge == 1 && !visitedMap[edgeInStr.Label]{
                    fmt.Println(g.InverseVertexTracker[j].Label)
					sn.Enqueue(edgeInStr)
                    visitedMap[edgeInStr.Label] = true
				}
			}
			break
		}
	}
}


type Queue struct {
	Data *GraphNode
	Next *Queue
}

type QueueUtil struct {
	Head *Queue
	Tail *Queue
}

func (q *QueueUtil) Enqueue(data *GraphNode){
	if q.Head == nil{
		q.Head = &Queue{Data:data, Next:nil}
		q.Tail = q.Head
		return
	}
	newVertex := &Queue{Data:data, Next:nil}
	q.Tail.Next = newVertex
	q.Tail = newVertex
}

func (q *QueueUtil) Dequeue()*GraphNode{
	if q.Head == nil{
		return nil
	}
	temp := q.Head
	if q.Head.Next != nil{
		q.Head = q.Head.Next
	}else {
		q.Head = nil
	}
	return temp.Data
}

func (q *QueueUtil) IsEmpty()bool{
	return q.Head == nil
}

