package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/goware/urlx"
)

func main() {
	// set ENV variable
	_ = os.Setenv("MY_SERVER", "https://some.server.com:9092/connect")

	// accept user input and use default value form ENV
	serverHost := flag.String("server", os.Getenv("MY_SERVER"), "a server host:ip/path")
	flag.Parse()

	// print supplied parameter
	fmt.Println("Server:", *serverHost)

	// parse server host
	url, err := urlx.Parse(*serverHost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scheme:", url.Scheme)
	fmt.Println("Hostname:", url.Hostname())
	fmt.Println("Port:", url.Port())
}
