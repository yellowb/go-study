package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	fmt.Printf("[Parent] started to invoke goroutine\n")
	go childFunc(1, 2, ch)
	fmt.Printf("[Parent] waiting for goroutine data...\n")
	result := <-ch
	fmt.Printf("[Parent] got result from goroutine with result = %d\n", result)
}

func childFunc(a int, b int, ch chan int) {
	fmt.Printf("[Child] got invoked with a = %d, b = %d\n", a, b)
	time.Sleep(3 * time.Second)
	fmt.Printf("[Child] calculation finished with result = %d\n", a+b)
	ch <- a + b
}
