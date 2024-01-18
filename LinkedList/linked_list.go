package main

import "fmt"

type linkedList struct {
	head *node
}

type node struct {
	value int
	next  *node
}

func (list *linkedList) addNodeAtTheEnd(value int) {

	// Cria um novo nó
	newNode := &node{value: value}

	// Se a lista estiver vazia, esse nó entra como primeiro
	if list.head == nil {
		list.head = newNode
		return
	} else { // Caso contrário, vamos até o fim e colocá-lo como último
		currentNode := list.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}

func (list *linkedList) addNodeAtTheBeggining(value int) {

	// Cria um novo nó
	newNode := &node{
		value: value,
		next:  list.head,
	}

	// E adiciona o nó como primeiro
	list.head = newNode
}

func (list linkedList) displayList() {
	fmt.Println("Lista Atual")

	// Passar por cada nó printando o resultado.
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func (list *linkedList) clearList() {
	list.head = nil
}

func (list linkedList) length() {
	length := 0
	current := list.head

	for current != nil {
		length += 1
		current = current.next
	}

	fmt.Printf("A lista possui %d nós.", length)
}

func (list *linkedList) deleteNodeByValue(value int) {

	// Se for o primeiro nó, apenas apague ele.
	if list.head != nil && list.head.value == value {
		list.head = list.head.next
		return
	}

	var previousNode *node
	currentNode := list.head

	// Caso contrário, vamos ir passando pela lista até achar o valor.
	for currentNode != nil && currentNode.value != value {
		previousNode = currentNode
		currentNode = currentNode.next
	}

	// E fazer com que o nó anterior aponte para o próximo do nó que será deletado.
	if currentNode != nil {
		previousNode.next = currentNode.next
	}
}

func main() {
	myList := linkedList{
		head: nil,
	}

	myList.addNodeAtTheEnd(1)
	myList.addNodeAtTheEnd(2)
	myList.displayList()
	myList.length()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Limpando a Lista...")
	myList.clearList()
	myList.displayList()

	fmt.Println("")
	fmt.Println("Adicionando novos números...")
	myList.addNodeAtTheEnd(1)
	myList.addNodeAtTheEnd(2)
	myList.addNodeAtTheBeggining(3)
	myList.addNodeAtTheBeggining(4)
	myList.addNodeAtTheBeggining(5)
	myList.displayList()

	fmt.Println("")
	fmt.Println("Deletando um valor...")
	myList.deleteNodeByValue(3)
	myList.displayList()
}
