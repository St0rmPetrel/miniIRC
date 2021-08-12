package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Users struct {
	mu     sync.Mutex
	name   map[string]net.Conn
	mirror map[net.Conn]string
}

func NewUsers() *Users {
	return &Users{name: make(map[string]net.Conn),
		mirror: make(map[net.Conn]string)}
}

func (users *Users) Add(conn net.Conn) error {
	var user_name string

	if _, err := fmt.Fscanln(conn, &user_name); err != nil {
		log.Printf("Connection Error: client introduce Error: %s\n", err)
		return &ConnectionErr{}
	}
	log.Printf("User: \"%s\" try to connect...\n", user_name)

	users.mu.Lock()
	if _, ok := users.name[user_name]; ok {
		users.mu.Unlock()
		fmt.Fprintln(conn, "false")
		log.Printf("User: \"%s\" connection refuse\n", user_name)
		return &NameExistErr{}
	}
	users.name[user_name] = conn
	users.mirror[conn] = user_name
	users.mu.Unlock()

	if _, err := fmt.Fprintln(conn, "true"); err != nil {
		log.Printf("Connection Error: client accept Error: %s\n", err)
		return &ConnectionErr{}
	}
	log.Printf("User: \"%s\" connection accept\n", user_name)
	return nil
}

func (users *Users) Delete(conn net.Conn) {
	log.Printf("Delete user \"%s\"\n", users.mirror[conn])
	delete(users.name, users.mirror[conn])
	delete(users.mirror, conn)
}

type ConnectionErr struct {
}

func (err *ConnectionErr) Error() string {
	return "Connection error with client"
}

type NameExistErr struct {
}

func (err *NameExistErr) Error() string {
	return "Name is alreade exist error"
}
