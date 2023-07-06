package domain

type Node struct {
	Next *Node
	Data any
}

type List struct {
	Head *Node
}
