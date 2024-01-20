package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) add(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) addListToStack(values ...int) {
	s.items = append(s.items, values...)
}

func (s *Stack) isEmpty() bool {
	if len(s.items) > 0 {
		return false
	}
	return true
}

func (s *Stack) size() int {
	return len(s.items)
}

func (s *Stack) clear() {
	s.items = []int{}
}

func (s *Stack) pop() {
	if s.isEmpty() {
		fmt.Println("Impossível remover items, a pilha está vazia!")
	} else {
		indiceRemocao := len(s.items) - 1 // Índice do Último item da lista
		s.items = s.items[:indiceRemocao]
	}
}

func (s *Stack) showStack() {
	if s.isEmpty() {
		fmt.Println("A Pilha está vazia!")
	} else {
		fmt.Printf("Início -> ")
		for _, value := range s.items {
			fmt.Printf("%v -> ", value)
		}
		fmt.Printf("Fim")
		fmt.Println("")
	}
}

func main() {
	myStack := Stack{}

	fmt.Println("Testando Adição...")
	myStack.add(3)
	myStack.add(7)
	myStack.add(18)
	myStack.add(11)
	myStack.addListToStack(22, 39, 10, 5)
	myStack.showStack()
	fmt.Printf("O tamanho da pilha é %v\n", myStack.size())

	fmt.Println("")
	fmt.Println("Removendo Itens...")
	myStack.pop()
	myStack.pop()
	myStack.pop()
	myStack.pop()
	myStack.pop()
	myStack.pop()
	myStack.showStack()
	fmt.Printf("O tamanho da pilha é %v\n", myStack.size())

	fmt.Println("")
	fmt.Println("Limpando Fila...")
	myStack.clear()
	myStack.pop()
	fmt.Printf("O tamanho da fila é %v\n", myStack.size())
}
