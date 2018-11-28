package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileName := "D:/lines.txt"
	counts := make(map[string]int)
	// Read file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "dup3 error: %s", err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}
	for line, count := range counts {
		fmt.Printf("%s -> %d\n", line, count)
	}
}
