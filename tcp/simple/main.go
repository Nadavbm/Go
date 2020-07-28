// simple tcp server
// run and connect via port 13000
package main

import (
	"io"
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
			log.Println(err)
		}

		io.WriteString(conn, "\nHello World\n\n")

		conn.Close()
	}
}
