package main

import (
	"fmt"
)

type Node struct {
	x int
	y int
}

func main() {
	node := Node{x: 1}
	node.y = 10

	p := &node
	p.y = 999

	fmt.Println(node)
	fmt.Println(p.y)
}
