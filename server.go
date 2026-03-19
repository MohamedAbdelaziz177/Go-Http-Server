package main

import (
	"fmt"
	"net"
)

type Server struct {
	router  *Router
	address string
}

func NewServer(address string) *Server {
	return &Server{
		router:  NewRouter(),
		address: address,
	}
}

func (s *Server) ListenAndStart() error {
	listner, err := net.Listen("tcp", s.address)

	if err != nil {
		return fmt.Errorf("Error Starting The Server: %s", err)
	}

	defer listner.Close()

	fmt.Printf("Server is running on %s\n", s.address)

	for {
		conn, err := listner.Accept()

		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) error {
	defer conn.Close()

	request, err := ParseRequest(conn)

	if err != nil {
		return fmt.Errorf("Error Parsing The Request: %s", err)
	}

	fmt.Printf("Request: %s %s\n", request.Method, request.Path)

	response := NewResponse(200, make(map[string]string), nil)

	s.router.ServeRequest(request, response)

	return response.Send(conn)
}
