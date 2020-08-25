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
	fromnodeVertex := g.VertexTracker[node1]
	tonodeVertex := g.VertexTracker[node2]
	if fromnodeVertex == nil || tonodeVertex == nil{
		return
	}
	g.matrix[fromnodeVertex.Position][tonodeVertex.Position] = 1
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
	sn := StackUtil{}
	visitedMap[startingNode.Label] = true
	sn.Push(startingNode)
	fmt.Println(startingNode.Label)
	for !sn.IsEmpty(){
		element := sn.Pop()
		elementPosition := g.VertexTracker[element.Label].Position
		for i, edges := range g.matrix{
			if i != elementPosition{
				continue
			}
			for j, edge := range edges{
				edgeInStr := g.InverseVertexTracker[j]
				if edge == 1 && !visitedMap[edgeInStr.Label]{
                    fmt.Println(g.InverseVertexTracker[j].Label)
					sn.Push(edgeInStr)
                    visitedMap[edgeInStr.Label] = true
				}
			}
		}
	}
}



type StackNode struct {
	Data *GraphNode
	Next *StackNode
}

type StackUtil struct {
	Head *StackNode
}

func (stk *StackUtil) Push(data *GraphNode){
	if data == nil{
		return
	}
	if stk.Head == nil{
		stk.Head = &StackNode{Data:data, Next:nil}
		return
	}
	newNode := &StackNode{Data:data, Next:stk.Head}
	stk.Head = newNode
}

func (stk *StackUtil) Pop()*GraphNode{
	if stk.Head == nil{
		return nil
	}

	temp := stk.Head
	if stk.Head.Next != nil{
		stk.Head = stk.Head.Next
	}else{
		stk.Head = nil
	}
	return temp.Data
}

func (stk *StackUtil) IsEmpty()bool{
	return stk.Head == nil
}

func (stk *StackUtil) Peek()*GraphNode{
	if stk.Head == nil{
		return nil
	}
	return stk.Head.Data
}

