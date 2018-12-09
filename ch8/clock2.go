package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("[Server]started to accept new connection.")
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("[Server]accepted connection - %v\n", conn)
		go handleConnUsingGoroutine(conn)
	}
}

func handleConnUsingGoroutine(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			fmt.Printf("[Server]client error - %v\n", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
