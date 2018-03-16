package main

import "fmt"

func main() {
	// Full define
	var words [2]string
	words[0] = "abc"
	words[1] = "def"

	fmt.Println(words[0], words[1])
	fmt.Println(words)

	// Short define
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slice
	slice := primes[1:4]
	fmt.Println(slice)
	slice[0] = 999
	fmt.Println(slice)
	fmt.Println(primes)

	// nil slice
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
