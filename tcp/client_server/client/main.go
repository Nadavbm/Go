// client represent a client that can dial server
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var addr = "localhost:12345"

func main() {
	fmt.Println("what you are going to say: ")
	for {
		text := textScanner()

		dialHandler(addr, text)
	}
}

func dialHandler(address, text string) {
	dial, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln("FATAL:", err)
	}
	defer dial.Close()

	fmt.Fprintln(dial, text)
}

func textScanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return scanner.Text()
	}
	return "end"
}
