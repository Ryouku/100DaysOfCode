package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/goware/urlx"
)

func main() {
	serverHost := flag.String("server", "", "a server host:ip/path")
	flag.Parse()

	fmt.Println("Server:", *serverHost)

	url, err := urlx.Parse(*serverHost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scheme:", url.Scheme)
	fmt.Println("Hostname:", url.Hostname())
	fmt.Println("Port:", url.Port())
}
