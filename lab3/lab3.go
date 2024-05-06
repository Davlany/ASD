package main

import (
	"fmt"
	lab3 "lab3/linked_list"
)

func main() {
	var (
		alphabetLinkedList lab3.LinkedList
		inputLinkedList    lab3.LinkedList
		resLinkedList      lab3.LinkedList
		input              string
	)

	alphabetLinkedList.PushMany('a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z')
	fmt.Scanln(&input)
	for _, v := range input {
		inputLinkedList.AddLastNode(v)
	}
	iterator := inputLinkedList.Iterator()
	vordNum := 0
	symbolNum := 0
	inWord := false
	skip := false
	for iterator.Next() {
		if iterator.CurrentNode.Data == rune(' ') && vordNum < 1 {
			vordNum++
		} else if iterator.CurrentNode.Data == rune(' ') && vordNum >= 1 {
			symbolNum = 0
			inWord = true
			skip = false
		} else {
			if skip {
				continue
			}
			if alphabetLinkedList.GetNodeByIndex(symbolNum).Data == iterator.CurrentNode.Data && vordNum >= 1 && inWord {
				resLinkedList.AddLastNode(iterator.CurrentNode.Data)
				symbolNum++
			} else if alphabetLinkedList.GetNodeByIndex(symbolNum).Data != iterator.CurrentNode.Data && vordNum >= 1 && inWord {
				resLinkedList.Clear()
				skip = true
			}
		}
	}
	resLinkedList.PrintList()

}
