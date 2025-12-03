# Real-Time Concurrent Chat System in Go

This project implements a real-time multi-client chat system using Go.\
It replaces the RPC-based approach with a concurrent TCP broadcasting
architecture as required in Assignment 04.

The system supports multiple clients communicating simultaneously, uses
goroutines for concurrency, maintains a synchronized list of active
clients, and broadcasts messages in real time without echoing messages
back to the sender.

------------------------------------------------------------------------

## Features

### Real-Time Message Broadcasting

-   Messages sent by a client are delivered immediately to all other
    connected clients.
-   The sender does not receive their own message.

### User Join and Leave Notifications

-   When a client connects, all existing clients are notified:
    `User [ClientX] joined.`
-   When a client disconnects or types `exit`, a leave notification is
    broadcasted.

### Concurrency and Synchronization

-   Each client connection is handled in a separate goroutine.
-   A `sync.Mutex` protects the shared client list to avoid race
    conditions.

### Simple Architecture

-   Uses only Go's standard library.
-   No RPC, frameworks, or external dependencies.

------------------------------------------------------------------------

## Project Structure

    .
    ├── server.go   # Chat server and broadcasting logic
    └── client.go   # Chat client program

------------------------------------------------------------------------

## How to Run

### 1. Start the Server

Open a terminal inside the project directory and run:

``` bash
go run server.go
```

Expected output:

    Chat server running on port 1234...

### 2. Start a Client

Open a new terminal window and run:

``` bash
go run client.go
```

A client ID (Client1, Client2, etc.) will be assigned automatically.

### 3. Start Additional Clients

Open more terminals and run:

``` bash
go run client.go
```

Each client will receive messages from others in real time.

------------------------------------------------------------------------

## Example Interaction

Client1 joins:\
`User [Client1] joined.`

Client2 joins:\
`User [Client2] joined.`

Client1 sends:\
`hello`

Client2 receives:\
`[Client1]: hello`

Client1 does not receive their own message.

------------------------------------------------------------------------

## Technical Overview

### Server

-   Maintains a map of connected clients.
-   Uses a mutex to synchronize shared access.
-   Runs `handleClient` in a separate goroutine for each new client.
-   Uses a `broadcast` function to send messages to all clients except
    the sender.

### Client

-   Uses one goroutine to listen for incoming messages.
-   Main thread handles user input.
-   Sends user text to the server until `exit` is entered.

------------------------------------------------------------------------

## Requirements Covered

-   Multi-client real-time communication\
-   Broadcasting without self-echo\
-   Concurrency using goroutines\
-   Shared memory protection using mutex\
-   User join and leave notifications\
-   TCP-based communication (not RPC)\
-   New GitHub repository exclusively for this assignment
