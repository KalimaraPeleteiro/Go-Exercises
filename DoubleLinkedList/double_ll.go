package main

import "fmt"

type Node struct {
	value        int
	nextNode     *Node
	previousNode *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (list *DoubleLinkedList) addNodeAtTheEnd(value int) {
	newNode := &Node{value: value, nextNode: nil, previousNode: nil}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.previousNode = list.tail
		list.tail.nextNode = newNode
		list.tail = newNode
	}
}

func (list *DoubleLinkedList) addNodeAtTheBeggining(value int) {
	newNode := &Node{value: value, nextNode: list.head, previousNode: nil}

	if list.head != nil {
		list.head.previousNode = newNode
	}

	list.head = newNode

	if list.tail == nil {
		list.tail = newNode
	}
}

func (list *DoubleLinkedList) deleteNodeByValue(value int) {
	var previousNode *Node

	currentNode := list.head

	// Vamos passar pela lista.
	for currentNode != nil {
		if currentNode.value == value { // Se encontrarmos o valor...
			if currentNode == list.head { // Caso 01: O nó a ser deletado é o primeiro.
				list.head = currentNode.nextNode
			}

			if currentNode == list.tail { // Caso 02: O nó a ser deletado é o último.
				list.tail = currentNode.previousNode
			}

			// Caso 03: O nó não é o primeiro nem o último.
			if previousNode != nil { // Atualize o Nó Anterior...
				previousNode.nextNode = currentNode.nextNode
			}

			if currentNode.nextNode != nil { // ... eo próximo nó.
				currentNode.nextNode.previousNode = currentNode.previousNode
			}

			return
		}

		previousNode = currentNode
		currentNode = currentNode.nextNode
	}
}

func (list *DoubleLinkedList) printList() {
	currentNode := list.head

	fmt.Println("Printando a Lista...")
	for currentNode != nil {
		fmt.Printf("Valor do Nó: %d\n", currentNode.value)
		currentNode = currentNode.nextNode
	}
	fmt.Println("Fim da Lista.")
}

func main() {
	myList := DoubleLinkedList{}
	myList.addNodeAtTheEnd(10)
	myList.addNodeAtTheEnd(20)
	myList.addNodeAtTheEnd(30)
	myList.addNodeAtTheEnd(40)
	myList.addNodeAtTheBeggining(50)
	myList.addNodeAtTheBeggining(60)
	myList.addNodeAtTheBeggining(70)

	myList.printList()

	fmt.Println("")
	fmt.Println("Apagando alguns valores...")
	myList.deleteNodeByValue(30)
	myList.deleteNodeByValue(60)
	myList.deleteNodeByValue(10)

	myList.printList()
}
