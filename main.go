package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	socks5 "github.com/armon/go-socks5"
)

const (
	defaultHost = "0.0.0.0"
	defaultPort = 1080
)

func main() {
	username := flag.String("username", "", "User login")
	password := flag.String("password", "", "User password")
	host := flag.String("host", defaultHost, "Host to listen")
	port := flag.Int("port", defaultPort, "Port to listen")
	flag.Parse()

	var conf socks5.Config
	switch {
	case *username != "" && *password != "":
		creds := map[string]string{*username: *password}
		conf.Credentials = socks5.StaticCredentials(creds)
	case *username == "" && *password == "":
		log.Println("WARNING: Anonymous users are enabled")
	default:
		log.Println("Username and password must be either blank or both not blank")
		os.Exit(1)
	}

	socks, err := socks5.New(&conf)
	if err != nil {
		log.Printf("Failed to make server: %v\n", err)
		os.Exit(1)
	}

	server := &Server{origin: socks}
	addr := fmt.Sprintf("%s:%d", *host, *port)

	log.Printf("Listening on %s\n", addr)
	if err := server.ListenAndServe("tcp", addr); err != nil {
		log.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
}
