package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Data  int
	Right *Node
	Left  *Node
}

type BinaryTree struct {
	RootNode    *Node
	CurrentNode *Node
}

func PrintTree(root *Node) {
	printNode(root, 0)
}

func Search(root *Node, target int) bool {
	if root == nil {
		return false
	}

	if root.Data == target {
		return true
	} else if target < root.Data {
		return Search(root.Left, target)
	} else {
		return Search(root.Right, target)
	}
}

func printNode(node *Node, level int) {
	if node == nil {
		return
	}

	printNode(node.Right, level+1)

	fmt.Println(strings.Repeat("   ", level), node.Data)

	printNode(node.Left, level+1)
}

func (bt *BinaryTree) AddNode(num int, rootNode ...*Node) {
	if rootNode == nil {
		bt.CurrentNode = bt.RootNode
	} else {
		bt.CurrentNode = rootNode[0]
	}
	if num > bt.CurrentNode.Data {
		if bt.CurrentNode.Right == nil {
			bt.CurrentNode.Right = &Node{
				Data: num,
			}
			return
		} else {
			bt.AddNode(num, bt.CurrentNode.Right)
		}
	} else {
		if bt.CurrentNode.Left == nil {
			bt.CurrentNode.Left = &Node{
				Data: num,
			}
			return
		} else {
			bt.AddNode(num, bt.CurrentNode.Left)
		}
	}
}

func FindMax(root *Node) (int, bool) {
	if root == nil {
		return 0, false
	}

	current := root
	for current.Right != nil {
		current = current.Right
	}

	return current.Data, true
}

func NewBinaryTree(num int) *BinaryTree {
	return &BinaryTree{
		RootNode: &Node{
			Data: num,
		},
	}
}

func main() {
	bt := NewBinaryTree(9)
	bt.AddNode(11)
	bt.AddNode(5)
	bt.AddNode(8)
	bt.AddNode(15)
	bt.AddNode(22)
	PrintTree(bt.RootNode)
	found := Search(bt.RootNode, 28)
	if found {
		fmt.Printf("Число %d знайдено в дереві.\n", 28)
	} else {
		fmt.Printf("Число %d не знайдено в дереві.\n", 28)
	}
	max, found := FindMax(bt.RootNode)
	if found {
		fmt.Println("Максимальный элемент в дереве:", max)
	} else {
		fmt.Println("Дерево пустое.")
	}
}
