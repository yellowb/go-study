package main

import "fmt"

func main() {
	var animal Animal
	dog := Dog{name: "Tom"}
	animal = &dog
	fmt.Println(animal.run())
}

type Animal interface {
	run() string
}

type Dog struct {
	name string
}

func (dog *Dog) run() string {
	return fmt.Sprintf("I am a running dog. My name is %s", dog.name)
}
