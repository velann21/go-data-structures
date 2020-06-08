package main

import "sync"

func main(){

}

// Node a single node that composes the tree
type Node struct {
	key   string
	left  *Node //left
	right *Node //right
}

type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func (bst *ItemBinarySearchTree) Insert(key string) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &Node{key,  nil, nil}
	if bst.root == nil{
		bst.root = n
	}else{
		bst.insertNode(bst.root, n)
	}
}

func (bst *ItemBinarySearchTree) insertNode(node, newNode *Node){
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			bst.insertNode(node.left, newNode)
		}
	}else {
		if node.right == nil {
			node.right = newNode
		} else {
			bst.insertNode(node.right, newNode)
		}
	}
}





