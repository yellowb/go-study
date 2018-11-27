package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "D:/lines.txt"
	counts := make(map[string]int)
	f, err := os.Open(fileName)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	} else {
		countLines(f, counts)
	}
	_ = f.Close()
	for line, count := range counts {
		fmt.Printf("%s -> %d\n", line, count)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
