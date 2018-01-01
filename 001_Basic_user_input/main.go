package main

import (
	"fmt"
	"log"
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

	// log event
	log.Println("Entered: " + server)

	fmt.Print("Enter server port: ")
	fmt.Scanln(&port)

	// need to convert int to string in order to get string concetenation
	sPort := strconv.Itoa(port)

	// log event
	log.Println("Entered: " + sPort)

	fmt.Print("Enter API path: ")
	fmt.Scanln(&apiPath)

	// log event
	log.Println("Entered: " + apiPath)

	// do final string concetenation and present the result
	fmt.Println("http://" + server + ":" + sPort + "/" + apiPath)
}
