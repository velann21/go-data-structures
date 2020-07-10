package main

import "fmt"

func main() {
	node := GraphItems{}
	nA := node.AddNode("A")
	nB := node.AddNode("B")
	nc := node.AddNode("C")
	nD := node.AddNode("D")
	nE := node.AddNode("E")
	nF := node.AddNode("F")

	node.AddEdges(nA, nB)
	node.AddEdges(nA, nc)
	node.AddEdges(nB, nE)
	node.AddEdges(nE, nF)
	node.AddEdges(nc, nD)
	node.AddEdges(nD, nF)

	node.PrintItems()

}

type GraphNode struct {
	Data string
}

type GraphItems struct {
	Items []*GraphNode
	Degrees map[*GraphNode][]*GraphNode
}

func (i *GraphItems) AddNode(data string)*GraphNode{
	gn := &GraphNode{Data:data}
	i.Items = append(i.Items, gn)
	return gn
}

func (i *GraphItems) AddEdges(node1 *GraphNode, node2 *GraphNode){
	if i.Degrees == nil{
		i.Degrees = make(map[*GraphNode][]*GraphNode)
	}
	i.Degrees[node1] = append(i.Degrees[node1], node2)
	i.Degrees[node2] = append(i.Degrees[node2], node1)
}

func (i *GraphItems) PrintItems(){
	for _, v := range i.Items{
		fmt.Print(v, "->")
		for _, v := range i.Degrees[v]{
			fmt.Print(v)
		}
		fmt.Println()
	}
}

func (i *GraphItems) Bfs(){

}
