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
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	if host == "" {
		host = defaultHost
	}
	port := defaultPort
	if value := os.Getenv("PORT"); value != "" {
		var err error
		port, err = strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Invalid PORT %q: %v", value, err)
		}
	}

	var conf socks5.Config
	switch {
	case username == "":
		log.Println("WARNING: Running in anonymous mode")
	case username != "" && password != "":
		creds := map[string]string{username: password}
		conf.Credentials = socks5.StaticCredentials(creds)
	default:
		log.Println("Password must not be blank when username is set")
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
