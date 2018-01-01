package main

import (
	"fmt"
	"strconv"
)

var (
	apiPath string
	server  string
	port    int
)

func main() {
	fmt.Print("Enter server address: ")
	fmt.Scanln(&server)

	fmt.Print("Enter server port: ")
	fmt.Scanln(&port)

	fmt.Print("Enter API path: ")
	fmt.Scanln(&apiPath)

	// need to convert int to string in order to get string concetenation
	sPort := strconv.Itoa(port)
	fmt.Println("http://" + server + ":" + sPort + "/" + apiPath)
}
