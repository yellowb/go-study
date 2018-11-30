package main

import "fmt"

type People struct {
	Name   string
	Age    int
	Friend *People
}

func main() {
	var tom People
	tom.Name = "Tom"
	tom.Age = 12

	// Pointer
	tomPtr := &tom
	fmt.Println(tomPtr.Age)

	// Below 2 cases are the same
	(*tomPtr).Age++
	tomPtr.Age++

	fmt.Println(tom)

	// Another people
	var mary People
	mary.Name = "Mary"
	mary.Age = 19
	mary.Friend = tomPtr
	mary.Friend.Age++ // can access

	fmt.Println(mary)
	fmt.Println(tom)
}