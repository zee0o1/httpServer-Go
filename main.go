package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("listen error %s", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("accept error %s", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("We are Here client is %s", conn.RemoteAddr())
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Buffer: %s\n", buf)
	_, err = conn.Write([]byte("HTTP/1.0 200 OK\nDate: Sat, 03 Oct 2026 13:29:45 GMT\nContent-Length: 13\nContent-Type: text/plain; charset=utf-8\nHello, \"/bar\""))
	if err != nil {
		log.Fatalf("write error %s", err)
	}
	err = conn.Close()
	if err != nil {
		log.Fatalf("connection close error %s", err)
	}
}
