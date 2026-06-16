package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080")

	// Goroutine to receive messages
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println("\nReceived: " + scanner.Text())
			fmt.Print("> ")
		}
	}()

	// Main loop to send messages
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Fprintln(conn, text)
	}
}
