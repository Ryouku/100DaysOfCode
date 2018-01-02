package main

import (
	"fmt"
	"log"

	"github.com/goware/urlx"
)

var (
	rawurl string
)

func main() {
	fmt.Print("Enter service URL in the FQDN:port/path format: ")

	_, err := fmt.Scanln(&rawurl)
	checkErr(err)

	url, err := urlx.Parse(rawurl)
	checkErr(err)

	fmt.Println("Parsed data follows...")

	fmt.Println("Schema: " + url.Scheme)
	fmt.Println("Host: " + url.Hostname())
	fmt.Println("Port: " + url.Port())
	fmt.Println("Path: " + url.Path)
}

// error checker helper
func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
