package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (tree *BinaryTree) insert(value int) {
	tree.Root = insertNode(tree.Root, value)
}

func insertNode(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}
	if value < node.Value {
		node.Left = insertNode(node.Left, value)
	} else {
		node.Right = insertNode(node.Right, value)
	}
	return node
}

func (tree *BinaryTree) Print() {
	printInOrder(tree.Root)
}

func printInOrder(node *Node) {
	if node != nil {
		printInOrder(node.Left)
		fmt.Print(node.Value, " ")
		printInOrder(node.Right)
	}
}

func (tree *BinaryTree) Search(value int) bool {
	return searchNode(tree.Root, value)
}

func searchNode(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	}
	if value < node.Value {
		return searchNode(node.Left, value)
	}
	return searchNode(node.Right, value)
}

func (tree *BinaryTree) Delete(value int) {
	tree.Root = deleteNode(tree.Root, value)
}

func deleteNode(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		minNode := findMin(node.Right)
		node.Value = minNode.Value
		node.Right = deleteNode(node.Right, minNode.Value)
	}
	return node
}

func findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func main() {

	tree := &BinaryTree{}
	tree.insert(12)
	tree.insert(5)
	tree.insert(3)
	tree.insert(1)
	tree.insert(11)
	tree.insert(9)
	tree.insert(6)
	tree.insert(8)
	tree.insert(7)
	tree.Print()
	fmt.Println("\n Search 11:", tree.Search(11))
	tree.Delete(9)
	tree.Print()
}
