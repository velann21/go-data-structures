package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {

	node := NodeItems3{}
	nA := node.AddNode(Node3{Item:"A"})
	nB := node.AddNode(Node3{Item:"B"})
	nC := node.AddNode(Node3{Item:"C"})
	nD := node.AddNode(Node3{Item:"D"})
	nE := node.AddNode(Node3{Item:"E"})
	nF := node.AddNode(Node3{Item:"F"})
	nG := node.AddNode(Node3{Item:"G"})

	node.AddEdges(nA, nB, 25)
	node.AddEdges(nB, nA, 25)


	node.AddEdges(nA, nC, 6)
	node.AddEdges(nC, nA, 6)

	node.AddEdges(nC, nD, 18)
	node.AddEdges(nD, nC, 18)

	node.AddEdges(nD, nE, 12)
	node.AddEdges(nE, nD, 12)

	node.AddEdges(nE, nF, 10)
	node.AddEdges(nF, nE, 10)

	node.AddEdges(nF, nB, 9)
	node.AddEdges(nB, nF, 9)

	node.AddEdges(nB, nG, 10)
	node.AddEdges(nG, nB, 10)

	node.AddEdges(nD, nG, 16)
	node.AddEdges(nG, nD, 16)

	node.PrintAllOutDegrees()

	nearBy := node.IntializeMap([]Node3{nA, nB, nC, nD, nE, nF, nG})

	spanningTree := [6][2]Node3{}


	f, s := node.FindMinimumAmong(nearBy, &spanningTree)
	fmt.Println(f, s)

	fmt.Println(nearBy)



}


type Node3 struct {
	Item string
	Weight int
}

type NodeItems3 struct {
	mu sync.Mutex
	Items []Node3
	OutDegrees map[Node3][]Node3
}

func (node *NodeItems3) AddNode(data Node3) Node3{
	node.mu.Lock()
	node.Items = append(node.Items, data)
	node.mu.Unlock()
	return data
}

func (node *NodeItems3) AddEdges(n1 Node3, n2 Node3, wght int){
	node.mu.Lock()
	if node.OutDegrees == nil{
		node.OutDegrees = make(map[Node3][]Node3)
	}
	n2.Weight = wght
	node.OutDegrees[n1] = append(node.OutDegrees[n1], n2)
	node.mu.Unlock()
}

func (node *NodeItems3) PrintAllOutDegrees(){
	for _, v := range node.Items{
		fmt.Print(v.Item, "->")
		for _, v1 := range node.OutDegrees[v]{

			fmt.Print(v1.Item, " ", v1.Weight)
			fmt.Print(", ")
		}
		fmt.Println()
	}
}

func (node *NodeItems3) FindMinimumAmong(nearBy map[Node3]int, st *[6][2]Node3)(Node3, Node3){
	min := math.MaxInt8
	var first Node3
	var sec Node3
	fmt.Println(first, sec)
	for _, v := range node.Items{
		for _, v1 := range node.OutDegrees[v]{
			if v1.Weight < min{
				min = v1.Weight
				first = v
				sec = v1
			}
		}
	}

	nearBy[sec] = 0
	nearBy[first] = 0


	st[0][0] = first
	st[0][1] = sec


	fmt.Println(st[0][0])
	fmt.Println(st[0][1])

	for K, v := range nearBy{
		fmt.Print(K, "-> ")
		if v != 0{
			for _, v := range node.OutDegrees[K]{
				if v == st[0][0]{

				}else if v == st[0][1]{

				}

			}
			fmt.Println()
		}else{
			fmt.Println()
		}
	}
	return first, sec
}

func (node *NodeItems3) IntializeMap(mapVal []Node3)map[Node3]int{
	nearBy := make(map[Node3]int)
	for _, v := range  mapVal {
		nearBy[v] = math.MaxInt8
	}
	return nearBy
}




