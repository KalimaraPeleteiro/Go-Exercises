package main

import (
	"fmt"
	"io"
	"os"
)

// Estruturas
type TreeNode struct {
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

// Funções
func (tree *BinaryTree) insertNewNode(value int) *BinaryTree {

	// Se a árvore estiver vazia, vamos adicionar este como o primeiro nó.
	if tree.root == nil {
		tree.root = &TreeNode{value: value, leftNode: nil, rightNode: nil}
	} else {
		tree.root.insertNewNode(value) // Caso contrário, vamos chamar o insertNewNode() do TreeNode
	}

	return tree
}

// Insert do TreeNode
func (node *TreeNode) insertNewNode(value int) {

	// Se o nó estiver vazio, algo está errado. Pare a função.
	if node == nil {
		return
	} else if value < node.value { // Caso contrário verifique o valor. Se for menor que o nó atual...
		if node.leftNode == nil { // Adicione a esquerda, se ela já não estiver ocupada.
			node.leftNode = &TreeNode{value: value, leftNode: nil, rightNode: nil}
		} else {
			node.leftNode.insertNewNode(value) // Se estiver, passe para o próximo nó.
		}
	} else { // Caso contrário, adicione a direita.
		if node.rightNode == nil {
			node.rightNode = &TreeNode{value: value, leftNode: nil, rightNode: nil}
		} else {
			node.rightNode.insertNewNode(value) // Se estiver ocupado, passe para o próximo nó.
		}
	}
}

func printTree(writer io.Writer, node *TreeNode, numberOfSpaces int, label rune) {

	// Se o nó for vazio, pare.
	if node == nil {
		return
	}

	// Caso contrário, printe os espaços (para fazer a hierarquia)
	for i := 0; i < numberOfSpaces; i++ {
		fmt.Fprintf(writer, " ")
	}

	// E os valores. Use "M" para o nó principal (main).
	fmt.Fprintf(writer, "%c:%v\n", label, node.value)

	// E chame os outros nós recursivamente até o fim.
	printTree(writer, node.leftNode, numberOfSpaces+2, 'L')
	printTree(writer, node.rightNode, numberOfSpaces+2, 'R')
}

func main() {
	tree := &BinaryTree{}
	tree.insertNewNode(100).
		insertNewNode(-20).
		insertNewNode(-50).
		insertNewNode(-15).
		insertNewNode(-60).
		insertNewNode(50).
		insertNewNode(60).
		insertNewNode(55).
		insertNewNode(85).
		insertNewNode(15).
		insertNewNode(5).
		insertNewNode(-10)
	printTree(os.Stdout, tree.root, 0, 'M')
}
