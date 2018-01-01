package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
)

var (
	rawurl       string
	parsedRawURL parsedURL
)

type parsedURL struct {
	schema string
	host   string
	port   string // port in "net/http" is an string?
	uri    string
	origin string
}

func main() {
	fmt.Print("Enter service URL in the FQDN:port/path format: ")
	fmt.Scanln(&rawurl)

	parsedRawURL, err := handleInputURL(rawurl)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Parsed data follows...")

	fmt.Println("Schema: " + parsedRawURL.schema)
	fmt.Println("Host: " + parsedRawURL.host)
	fmt.Println("Port: " + parsedRawURL.port)
	fmt.Println("URI: " + parsedRawURL.uri)

	fmt.Println("Reassembled input: " + parsedRawURL.origin)
}

// hndles raw URL parsing and ensures that required parts are defined
// and delivered as an parsedURL struct
func handleInputURL(rawurl string) (parsed parsedURL, err error) {
	// parse URL
	s, err := url.Parse(rawurl)
	if err != nil {
		log.Fatal(err)
		return parsedURL{}, err
	}

	if !s.IsAbs() {
		log.Fatal("URL scheme not provided!")
		return parsedURL{}, errors.New("URL scheme not provided")
	}

	if s.Hostname() == "" {
		log.Fatal("Hostname is not recognised!")
		return parsedURL{}, errors.New("Hostname is not recognised")
	}

	var p = parsedURL{
		s.Scheme,
		s.Host,
		s.Port(),
		s.RequestURI(),
		s.String(),
	}

	return p, nil
}
