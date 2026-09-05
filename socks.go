package main

import (
	"context"
	"fmt"
	"log"
	"net"

	socks5 "github.com/armon/go-socks5"
)

// Server is a wrapper around go-socks5, that counts incoming requests.
type Server struct {
	listener net.Listener
	origin   *socks5.Server
}

// ListenAndServe is used to create a listener and serve on it.
func (s *Server) ListenAndServe(ctx context.Context, network, addr string) error {
	var lc net.ListenConfig
	var err error
	s.listener, err = lc.Listen(ctx, network, addr)
	if err != nil {
		return fmt.Errorf("listen on %s: %w", addr, err)
	}

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			return fmt.Errorf("accept connection: %w", err)
		}
		go func() {
			if err := s.origin.ServeConn(conn); err != nil {
				log.Printf("Failed to serve connection: %v", err)
			}
		}()
	}
}

// Stop closes the listener and stops the server.
func (s *Server) Stop() error {
	return s.listener.Close() //nolint:wrapcheck
}
