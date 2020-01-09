package main

import "fmt"

type node struct{
	value int
	next *node
}

func main() {
	start := createList(1)
	second := start.addNode(2)
	fmt.Println(start.value)
	fmt.Println(start.next.value)
	third := start.addNode(3)
	fmt.Println(start.next.value)
	fmt.Println(third.next.value)
	fmt.Println(second.next.value)
}

func createList(val int) node {
	newNode := node{value: 1, next: nil}
	newNode.next = &newNode
	return newNode
}

func (n *node) addNode(val int) node {
	newNode := node{value: val, next: n.next}
	n.next = &newNode
	return newNode
}
