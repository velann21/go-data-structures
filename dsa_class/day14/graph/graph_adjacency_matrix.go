package main

import "fmt"

func main() {
	graph := Graph{}
	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")

	fmt.Println(graph.tracker)

	graph.AddEdges("A", "B")
	graph.AddEdges("B", "D")
	graph.AddEdges("A", "C")
	graph.AddEdges("C", "D")

	graph.PrintGraph()



}

type Vertex struct {
	Label string
	Position int
}

type Graph struct {
	matrix [4][4]int
	tracker map[string]Vertex
	Inversetracker map[int]Vertex
	PositionTracker int
}

func (g *Graph) AddNode(node string){
	if g.tracker == nil{
		g.tracker = make(map[string]Vertex)
		g.Inversetracker = make(map[int]Vertex)
	}
	newVertex := Vertex{Label:node, Position:g.PositionTracker}
	g.Inversetracker[g.PositionTracker] = newVertex
	g.PositionTracker = g.PositionTracker+1
    g.tracker[node] = newVertex
}

func (g *Graph) AddEdges(from string, to string){
	if g.tracker == nil{
		return
	}
	fromPosition := g.tracker[from]
	toPosition := g.tracker[to]
	g.matrix[fromPosition.Position][toPosition.Position] = 1
}

func (g *Graph) PrintGraph(){
	for i, v := range g.matrix{
		fmt.Print(g.Inversetracker[i], "---->")
		for j , loopInner := range v{
			if loopInner == 1{
				fmt.Print(g.Inversetracker[j],",")
			}
		}
		fmt.Println()
	}
}
