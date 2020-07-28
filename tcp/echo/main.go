// this tcp server will echo every input line
// connect via port 13000
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":13000")
	if err != nil {
		log.Panic(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Fprintf(conn, "say what??!\t\t%s\n", txt)
	}
	defer conn.Close()
}
