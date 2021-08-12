package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/St0rmPetrel/miniIRC/client/userinterface"
)

func main() {
	user_name, addr, err_init := flag_init()
	if err_init != nil {
		println(fmt.Sprintf("Error: %s\n", err_init))
		return
	}
	conn, err_conn := net.Dial("tcp", addr)
	if err_conn != nil {
		println(fmt.Sprintf("Error: can't connect to server: %s: %s\n",
			addr, err_conn))
		return
	}
	if is_unique_name := check_name(user_name, conn); !is_unique_name {
		println(fmt.Sprintf("Error: name: \"%s\" is exist", user_name))
		return
	}
	session(conn)
}

func session(conn net.Conn) {
	clientEvents := make(chan string, 1)
	serverEvents := make(chan string, 10)
	quit := make(chan int, 1)
	go receive_msg(conn, serverEvents, quit)
	go send_msg(conn, clientEvents, quit)
	userinterface.Userinterface(clientEvents, serverEvents, quit)
	println("Connection with server is lost")
}

func receive_msg(conn net.Conn, serverEvents chan string, quit chan int) {
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			quit <- 1
			return
		}
		serverEvents <- msg
	}
}

func send_msg(conn net.Conn, clientEvents chan string, quit chan int) {
	for {
		msg := <-clientEvents
		if _, err := fmt.Fprintln(conn, msg); err != nil {
			quit <- 1
			return
		}
	}
}

func check_name(name string, conn net.Conn) (is_unique_name bool) {
	fmt.Fprintln(conn, name)
	fmt.Fscanln(conn, &is_unique_name)
	return is_unique_name
}
