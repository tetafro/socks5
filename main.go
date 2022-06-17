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
	anon := flag.Bool("anon", false, "Anonymous proxy")
	host := flag.String("host", defaultHost, "Host to listen")
	port := flag.Int("port", defaultPort, "Port to listen")
	flag.Parse()

	var conf socks5.Config
	switch {
	case *anon:
		log.Println("WARNING: Running in anonymous mode")
	case *username != "" && *password != "":
		creds := map[string]string{*username: *password}
		conf.Credentials = socks5.StaticCredentials(creds)
	default:
		log.Println("Username and password must not be blank")
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
