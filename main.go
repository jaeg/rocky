package main

import (
	"flag"
	"fmt"
	"net"
)

var target = flag.String("target", "localhost:6379", "Target address to forward traffic to.")
var port = flag.String("port", ":9999", "Port to run the reverse proxy on")

func main() {
	flag.Parse()
	ln, err := net.Listen("tcp", *port)
	if err != nil {
		fmt.Println(err)
		return
	}

	targetConn, err := net.Dial("tcp", *target)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Println("New connection")
		go handleToTarget(conn, targetConn)
		go handleFromTarget(conn, targetConn)
	}
}

func handleToTarget(conn net.Conn, targetConn net.Conn) {
	for {
		buf := make([]byte, 1)
		_, err := conn.Read(buf)
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
