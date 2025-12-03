package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

type Client struct {
	ID   string
	Conn net.Conn
}

type Server struct {
	mu      sync.Mutex
	clients map[string]Client
}

func NewServer() *Server {
	return &Server{
		clients: make(map[string]Client),
	}
}

// Broadcast message to all clients except sender
func (s *Server) broadcast(senderID string, msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, client := range s.clients {
		if id == senderID {
			continue // no self-echo
		}
		fmt.Fprintf(client.Conn, "%s\n", msg)
	}
}

func (s *Server) handleClient(client Client) {
	reader := bufio.NewScanner(client.Conn)

	// Notify others
	joinMsg := fmt.Sprintf("User [%s] joined.", client.ID)
	fmt.Println(joinMsg)
	s.broadcast(client.ID, joinMsg)

	for reader.Scan() {
		text := reader.Text()
		if text == "exit" {
			break
		}

		msg := fmt.Sprintf("[%s]: %s", client.ID, text)
		fmt.Println(msg)
		s.broadcast(client.ID, msg)
	}

	// Remove client on exit
	s.mu.Lock()
	delete(s.clients, client.ID)
	s.mu.Unlock()

	leaveMsg := fmt.Sprintf("User [%s] left.", client.ID)
	fmt.Println(leaveMsg)
	s.broadcast(client.ID, leaveMsg)

	client.Conn.Close()
}

func main() {
	server := NewServer()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Chat server running on port 1234...")

	idCounter := 1

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		clientID := fmt.Sprintf("Client%d", idCounter)
		idCounter++

		client := Client{
			ID:   clientID,
			Conn: conn,
		}

		server.mu.Lock()
		server.clients[clientID] = client
		server.mu.Unlock()

		go server.handleClient(client)
	}
}
