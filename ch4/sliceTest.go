package main

import "fmt"

func main() {
	array := [...]int{0, 1, 2, 3, 4, 5}
	s1 := array[:3]
	fmt.Println("s1.cap", cap(s1))
	fmt.Println("s1.len", len(s1))

	/* Case1: append one by one, will affect original array
	Output:
	[0 1 2 10 11 12 13]
	[0 1 2 10 11 12]
	*/
	//s1 = append(s1, 10)
	//s1 = append(s1, 11)
	//s1 = append(s1, 12)
	//s1 = append(s1, 13)

	/* Case2: append more than cap at one time, it will alloc a new low-layer array and not affect original array
	Output:
	[0 1 2 10 11 12 13]
	[0 1 2 3 4 5]
	*/
	s1 = append(s1, 10, 11, 12, 13)

	fmt.Println("s1 =", s1)
	fmt.Println("array =", array)

}
