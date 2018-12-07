package main

import (
	"../calc"
	"fmt"
)

func main() {
	a, b := 1, 2
	c := calc.Add(a, b)
	fmt.Println(c)
}
