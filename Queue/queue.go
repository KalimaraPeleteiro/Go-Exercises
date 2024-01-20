package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) add(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) addListToQueue(values ...int) {
	q.items = append(q.items, values...)
}

func (q *Queue) isEmpty() bool {
	if q.size() > 0 {
		return false
	}
	return true
}

func (q *Queue) size() int {
	return len(q.items)
}

func (q *Queue) clear() {
	q.items = []int{}
}

func (q *Queue) dequeue() {
	if q.isEmpty() {
		fmt.Println("A fila está vazia, não há o que deletar!")
	} else {
		q.items = q.items[1:]
	}
}

func (q *Queue) showQueue() {
	if q.isEmpty() {
		fmt.Println("A fila está vazia!")
	} else {
		fmt.Printf("Início -> ")
		for _, value := range q.items {
			fmt.Printf("%v -> ", value)
		}
		fmt.Printf("Fim")
		fmt.Println("")
	}
}

func main() {
	myQueue := Queue{}

	fmt.Println("Testando Adição...")
	myQueue.add(3)
	myQueue.add(7)
	myQueue.add(1)
	myQueue.addListToQueue(4, 13, 22, 11)
	myQueue.showQueue()
	fmt.Printf("O tamanho da fila é %v\n", myQueue.size())

	fmt.Println("")
	fmt.Println("Removendo Itens...")
	myQueue.dequeue()
	myQueue.dequeue()
	myQueue.dequeue()
	myQueue.showQueue()
	fmt.Printf("O tamanho da fila é %v\n", myQueue.size())

	fmt.Println("")
	fmt.Println("Limpando Fila...")
	myQueue.clear()
	myQueue.dequeue()
	fmt.Printf("O tamanho da fila é %v\n", myQueue.size())
}
