package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalln("FATAL:", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("ERROR:", err)
			continue
		}
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Println(txt)
	}
	defer conn.Close()
}
