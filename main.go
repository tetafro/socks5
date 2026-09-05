// socks5 is a SOCKS5 proxy server with password authentication.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	socks5 "github.com/armon/go-socks5"
)

const (
	defaultHost = "0.0.0.0"
	defaultPort = 1080
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

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
		log.Printf("Failed to make server: %v", err)
		os.Exit(1)
	}

	server := &Server{origin: socks}
	addr := fmt.Sprintf("%s:%d", host, port)
	go func() {
		<-ctx.Done()
		if err := server.Stop(); err != nil {
			log.Printf("Failed to stop server: %v", err)
			os.Exit(1)
		}
	}()

	log.Printf("Listening on %s", addr)
	if err := server.ListenAndServe(ctx, "tcp", addr); err != nil && ctx.Err() == nil {
		log.Printf("Failed to start server: %v", err)
		os.Exit(1)
	}
	log.Print("Shutdown gracefully")
}
