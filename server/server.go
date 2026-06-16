package main

import (
	"bufio"
	"fmt"
	"net"
)

var clients = make(map[net.Conn]bool)

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	fmt.Println("Server started on :8080")

	for {
		conn, _ := ln.Accept()
		clients[conn] = true
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		for client := range clients {
			if client != conn { // Send to the other person
				fmt.Fprintln(client, msg)
			}
		}
	}
	delete(clients, conn)
	conn.Close()
}
