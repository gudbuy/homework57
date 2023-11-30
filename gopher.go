package main

import "fmt"

type Node struct {
	value int
	next *Node
}

func (n1 Node) SummNode(n2 Node) Node {
	var one int
	var rootNode Node
	var node = &rootNode
	p1, p2 := &n1, &n2
	for {
		var summ int
		if p1 != nil {summ += p1.value}
		if p2 != nil {summ += p2.value}
		summ += one
		one = 0
		node.value = summ % 10
		if summ >= 10 {
			one = 1
		}
		if p1.next == nil && p2.next == nil && one == 0 {
			return rootNode
		}
		node.next = new(Node)
		node = node.next

		if p1.next != nil {p1 = p1.next} else {p1 = new(Node)}
		if p2.next != nil {p2 = p2.next} else {p2 = new(Node)}
	}
}

func main() {
	n1 := Node{1, &Node{2, &Node{3, &Node{5, nil}}}}
	n2 := Node{9, &Node{4, &Node{3, nil}}}
	n3 := n1.SummNode(n2)
	n3.Print()
}

func (n Node) Print() {
	for p := &n; p != nil; p = p.next {
		fmt.Print(p.value)
	}
}