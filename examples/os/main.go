package main

import (
	"fmt"
	"os/exec"
)

func checkFilesystem() {
	out, err := exec.Command("df", "-h").Output()
	if err != nil {
		fmt.Println("%s", err)
	}
	output := string(out[:])
	fmt.Println(output)
}

func listDir() {
	out, err := exec.Command("ls", "-lahs").Output()
	if err != nil {
		fmt.Println("%s", err)
	}
	fmt.Println("List directory")
	output := string(out[:])
	fmt.Println(output)
}

func main() {
	checkFilesystem()
	listDir()
}
