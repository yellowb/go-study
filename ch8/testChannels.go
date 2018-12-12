package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int)
	log.Printf("[Parent] started to invoke goroutine\n")
	go childFunc(1, 2, ch)
	go childFunc(2, 3, ch)
	log.Printf("[Parent] waiting for goroutine data...\n")
	result := <-ch
	log.Printf("[Parent] got result from goroutine with result = %d\n", result)
	result = <-ch
	log.Printf("[Parent] got result from goroutine with result = %d\n", result)
}

func childFunc(a int, b int, ch chan int) {
	log.Printf("[Child] got invoked with a = %d, b = %d\n", a, b)
	time.Sleep(time.Duration(a+b) * time.Second)
	log.Printf("[Child] calculation finished with result = %d\n", a+b)
	ch <- a + b
}
