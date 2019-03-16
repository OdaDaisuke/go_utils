package node

import "fmt"

type Node struct {
	Value string
	Prev *Node
	Next *Node
}

func Add(original, newNode *Node) *Node {
	original.Next = newNode
	newNode.Prev = original
	return newNode
}

func PrintAll(n *Node) {
	curNode := n

	for {
		fmt.Printf("Value: %s, Next: %p\n", curNode.Value, &curNode.Next)
		curNode = curNode.Next
		if curNode.Next == nil {
			break
		}
	}
}