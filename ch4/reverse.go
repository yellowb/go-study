package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	reverseSlice(array[:])
	reverseSlice(nil)
	fmt.Println(array)
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}