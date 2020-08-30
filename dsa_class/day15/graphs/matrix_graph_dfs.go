package main

import "fmt"

func main() {
	gt := GraphTraverser{}
	gt.AddNode("A")
	gt.AddNode("B")
	gt.AddNode("C")
	gt.AddNode("D")

	gt.AddEdges("A", "B")
	gt.AddEdges("B", "C")
	gt.AddEdges("C", "D")
	gt.AddEdges("A", "D")
	gt.AddEdges("D", "A")

	gt.PrintGraph()
	gt.DFS("A")

}

type Vertex struct {
	Label string
	Position int
}

type GraphTraverser struct {
	matrix [4][4]int
	positionTracker int
	VertexPositionMap map[string]*Vertex
	PositionVertexMap map[int]*Vertex
}

func (g *GraphTraverser) AddNode(vertex string){
	if vertex == ""{
		return
	}

	if g.VertexPositionMap == nil{
		g.VertexPositionMap = make(map[string]*Vertex)
		g.PositionVertexMap = make(map[int]*Vertex)
	}

	newVertex := &Vertex{Label:vertex, Position:g.positionTracker}
	g.VertexPositionMap[vertex] = newVertex
	g.PositionVertexMap[g.positionTracker] = newVertex
	g.positionTracker = g.positionTracker+1
}

func (g *GraphTraverser) AddEdges(from string, to string){
	if from == "" || to == "" {
		return
	}
	fromPosition := g.VertexPositionMap[from]
	toPosition := g.VertexPositionMap[to]
	if fromPosition == nil || toPosition == nil{
		return
	}
	g.matrix[fromPosition.Position][toPosition.Position] = 1
}

func (g *GraphTraverser) PrintGraph(){
	for edge, edges := range g.matrix{
		fmt.Print(g.PositionVertexMap[edge], "--->")
		for position, edges := range edges{
			if edges == 1{
				fmt.Print(g.PositionVertexMap[position])
			}
		}
		fmt.Println()
	}
}

func (g *GraphTraverser) DFS(start string){
	if start == ""{
		return
	}
	visitedMap := map[string]bool{}
	stk := StackUtil{}
	stk.Push(g.VertexPositionMap[start])
	visitedMap[start] = true
	fmt.Println(start)

	for !stk.IsEmpty(){
		data := stk.Pop()

		for i , edges := range g.matrix{
			if data.Position != i{
				continue
			}

			for j, innerEdges := range edges{
				vertex := g.PositionVertexMap[j]
				if innerEdges == 1 && !visitedMap[vertex.Label]{
					stk.Push(g.PositionVertexMap[i])
					data = vertex
					fmt.Println(vertex.Label)
					visitedMap[vertex.Label] = true
					break
				}

			}
		}
	}
}

type Stack struct {
	Data *Vertex
	Next *Stack
}

type StackUtil struct {
	Head *Stack
}

func (stk *StackUtil) Push(data *Vertex){
	if data == nil{
		return
	}
	if stk.Head == nil{
		stk.Head = &Stack{Data:data, Next:nil}
		return
	}
	newNode := &Stack{Data:data, Next:stk.Head}
	stk.Head = newNode
}

func (stk *StackUtil) Pop()*Vertex{
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

func (stk *StackUtil) Peek()*Vertex{
	if stk.Head == nil{
		return nil
	}
	return stk.Head.Data

}
