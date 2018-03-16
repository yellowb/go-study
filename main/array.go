package main

import "fmt"

func main() {
	var words [2]string
	words[0] = "abc"
	words[1] = "def"

	fmt.Println(words[0], words[1])
	fmt.Println(words)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
