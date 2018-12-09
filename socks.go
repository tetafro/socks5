package main

import (
	"log"
	"net"

	socks5 "github.com/armon/go-socks5"
)

// logFreq sets frequency of statistics output.
const logFreq = 1000

// Server is a wrapper around go-socks5, that counts incoming requests.
type Server struct {
	origin  *socks5.Server
	counter int
}

// ListenAndServe is used to create a listener and serve on it.
func (s *Server) ListenAndServe(network, addr string) error {
	l, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	return s.Serve(l)
}

// Serve is used to serve connections from a listener.
func (s *Server) Serve(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		s.counter++
		if s.counter%logFreq == 0 {
			log.Printf("Served %d requests", s.counter)
		}
		go func() {
			if err := s.origin.ServeConn(conn); err != nil {
				log.Printf("Failed to serve connection: %v", err)
			}
		}()
	}
}
