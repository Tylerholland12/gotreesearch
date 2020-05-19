package main

func main() {
	
}

import (
	
	"fmt"
)



type Node struct {
	Key int
	Left *Node
	Right *Node
}

tree := &Node{6, nil, nil}
fmt.Print(tree.key)


func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}
	if n.Key < key {
		return n.Right.Search(key)
	} else if n.Key > key{
		return n.Left.Search(key)
	}
	return True
}

func (n *Node) Insert(key int) {
	if n.Key < key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	} else if n.Key > key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	}
}

func (n *Node) Delete(key int) *Node {
	if n.Key < key {
		n.Right = n.Right.Delete(key)
	} else if n.Key > key {
		n.Left = n.Left.Delete(key)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		min := n.Right.Min()
		n.Key = min
		n.Right = n.Right.Delete(min)
	}
	return n
}

