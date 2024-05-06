package lab3

import (
	"fmt"
	"log"
)

type Node struct {
	Data        rune
	NextNode    *Node
	PreviusNode *Node
}

type LinkedList struct {
	Head         *Node
	Tale         *Node
	NodeIterator *Iterator
	Lenth        int
}

type Iterator struct {
	CurrentNode *Node
	Index       int
}

func (ll *LinkedList) AddLastNode(char rune) {
	if ll.Lenth == 0 {
		node := &Node{
			Data:        char,
			NextNode:    nil,
			PreviusNode: nil,
		}
		ll.Head = node
		ll.Tale = node
		ll.Lenth++
	} else if ll.Lenth == 1 {
		node := &Node{
			Data:        char,
			NextNode:    nil,
			PreviusNode: ll.Head,
		}
		ll.Tale = node
		ll.Head.NextNode = node
		ll.Lenth++
	} else {
		node := &Node{
			Data:        char,
			PreviusNode: ll.Tale,
			NextNode:    nil,
		}
		ll.Tale.NextNode = node
		ll.Tale = node
		ll.Lenth++
	}

}

func (ll *LinkedList) AddNodeByIndex(char rune, index int) {
	if index == 0 && ll.Lenth == 0 {
		ll.AddLastNode(char)
		return
	}
	if index > ll.Lenth-1 {
		log.Fatal("Error: index out of range")
	}
	if index == 0 {
		node := &Node{
			Data:        char,
			PreviusNode: nil,
			NextNode:    ll.Head,
		}
		ll.Head.PreviusNode = node
		ll.Head = node
		ll.Lenth++
		return
	} else if index != ll.Lenth-1 {
		newNode := &Node{
			Data: char,
		}
		node := ll.Head
		for i := 0; i < ll.Lenth; i++ {
			if i == index {
				node.PreviusNode.NextNode = newNode
				newNode.PreviusNode = node.PreviusNode
				newNode.NextNode = node
				node.PreviusNode = newNode
				ll.Lenth++
				return
			}
			node = node.NextNode
		}
	} else {
		ll.AddLastNode(char)
	}
}

func (ll *LinkedList) PrintList() {
	node := ll.Head
	for i := 0; i < ll.Lenth; i++ {
		if node.NextNode == nil {
			fmt.Print(string(node.Data))
			break
		}
		fmt.Print(string(node.Data))
		node = node.NextNode
	}
}

func (ll *LinkedList) Clear() {
	node := ll.Tale
	for i := 0; i < ll.Lenth; i++ {
		node.Data = 0
		node = node.PreviusNode
		node.NextNode = nil
	}
}

func (ll *LinkedList) GetNodeByIndex(index int) *Node {
	node := ll.Head
	for i := 0; i < index; i++ {
		node = node.NextNode
	}
	return node
}

func (ll *LinkedList) PushMany(symbols ...rune) {
	for _, v := range symbols {
		ll.AddLastNode(v)
	}
}

func (ll *LinkedList) Iterator() *Iterator {
	return &Iterator{
		CurrentNode: ll.Head,
		Index:       0,
	}
}

func (i *Iterator) Next() bool {
	if i.Index == 0 {
		i.Index++
		return true
	}
	node := i.CurrentNode
	if node.NextNode == nil {
		return false
	}
	i.CurrentNode = node.NextNode
	return true
}
