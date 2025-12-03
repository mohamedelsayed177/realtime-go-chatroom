package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Cannot connect:", err)
		return
	}
	defer conn.Close()

	// Receiving goroutine
	go func() {
		serverReader := bufio.NewScanner(conn)
		for serverReader.Scan() {
			fmt.Println(serverReader.Text())
		}
		os.Exit(0)
	}()

	// Sending loop
	userReader := bufio.NewReader(os.Stdin)
	for {
		text, _ := userReader.ReadString('\n')
		fmt.Fprintf(conn, text)
	}
}
