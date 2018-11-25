package main

import "fmt"

func main() {
	// Full define
	var words [2]string
	words[0] = "abcd"
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

	// Slice append
	// https://www.zhihu.com/question/27161493
	s := []int{5}
	printSlice(s)
	s = append(s, 7)
	printSlice(s)
	s = append(s, 9)
	printSlice(s)
	x := append(s, 11)
	printSlice(x)
	y := append(s, 12)
	printSlice(y)
	fmt.Println(s, x, y)

	fmt.Println("Another...")

	s = []int{5, 7, 9}
	printSlice(s)
	x = append(s, 11)
	printSlice(x)
	y = append(s, 12)
	printSlice(y)
	fmt.Println(s, x, y)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
