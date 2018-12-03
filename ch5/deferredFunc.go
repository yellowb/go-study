package main

import (
	"log"
	"time"
)

func main() {
	slowOp()
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("Enter %s\n", msg)
	return func() {
		log.Printf("Exit %s (%s)\n", msg, time.Since(start))
	}
}

func slowOp() {
	log.Printf("Start slow op...")
	defer trace("HAHAHA")()
	time.Sleep(3 * time.Second)
	log.Printf("Finish slow op")
}
