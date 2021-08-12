package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	addr := ":8080"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chat server is UP, addr = \"%s\"\n", addr)
	users := NewUsers()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error: %s\n", err)
			continue
		}
		go handleConnection(conn, users)
	}
}

func handleConnection(conn net.Conn, users *Users) {
	defer conn.Close()
	if err := users.Add(conn); err != nil {
		log.Printf("User add Error: %s\n", err)
		return
	}
	defer users.Delete(conn)
	session(conn, *users)
}

func session(conn net.Conn, users Users) {
	for {
		req, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("User: \"%s\" connetion is lost\n", users.mirror[conn])
			break
		}
		dest, msg := parseRequese(req)
		if msg == "" || dest == "" {
			continue
		} else if dest == "all" {
			broadSend(msg, users, conn)
		} else {
			privateSend(msg, dest, users, conn)
		}
	}
}

func parseRequese(req string) (dest, msg string) {
	if !strings.HasPrefix(req, "@") {
		return "all", req[:len(req)-1]
	}
	if id := strings.Index(req, " "); id > 0 {
		return req[1:id], req[id+1 : len(req)-1]
	}
	return req[1 : len(req)-1], ""
}

func broadSend(msg string, users Users, conn net.Conn) {
	log.Printf("User \"%s\" send messege \"%s\" broadcast", users.mirror[conn],
		msg)
	for _, users_conn := range users.name {
		_, err := fmt.Fprintln(users_conn, users.mirror[conn]+": "+msg)
		if err != nil {
			log.Printf("Error: Bad connection to user:%s\n",
				users.mirror[users_conn])
			continue
		}
	}
}

func privateSend(msg string, dest string, users Users, conn net.Conn) {
	log.Printf("User \"%s\" send private messege \"%s\" to user \"%s\"",
		users.mirror[conn], msg, dest)
	log.Printf("Try to find user \"%s\"...", dest)
	if dest_conn, ok := users.name[dest]; ok {
		log.Printf("User \"%s\" is found", dest)
		if _, err := fmt.Fprintln(conn, users.mirror[conn]+
			"->"+dest+": "+msg); err != nil {
			log.Printf("Error: Bad connection to user:%s\n",
				users.mirror[conn])
			return
		}
		if _, err := fmt.Fprintln(dest_conn, users.mirror[conn]+
			"->"+dest+": "+msg); err != nil {
			log.Printf("Error: Bad connection to user:%s\n", dest)
			return
		}
		return
	}
	log.Printf("Error: User \"%s\" not found", dest)
}
