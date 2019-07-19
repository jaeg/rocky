package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":6380")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("New connection")
	targetConn, err := net.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	go handleToTarget(conn, targetConn)
	go handleFromTarget(conn, targetConn)

}

func handleToTarget(conn net.Conn, targetConn net.Conn) {
	for {
		buf := make([]byte, 1)
		_, err := conn.Read(buf)
		fmt.Println("To Target:", buf)
		if err != nil {
			fmt.Println("Error reading to target:", err.Error())
			return
		}

		targetConn.Write(buf)

	}
}

func handleFromTarget(conn net.Conn, targetConn net.Conn) {
	for {
		buf := make([]byte, 1)
		_, err := targetConn.Read(buf)
		//stringed := string(buf)
		//fmt.Println("From target", stringed+"<end>")
		if err != nil {
			fmt.Println("Error reading from target:", err.Error())
			return
		}
		conn.Write(buf)
	}
}
