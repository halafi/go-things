package server

import (
	"fmt"
	"log"
	"net/http"
)

// Server struct containing hostname and port
type Server struct {
	Hostname string `json:"hostname"`
	HTTPPort string `json:"httpPort"`
}

// NewServer creates new instance of server
func NewServer(host, port string) *Server {
	return &Server{
		Hostname: host,
		HTTPPort: port,
	}
}

// Run starts the server at specified host and port
func (s *Server) Run(h http.Handler) {
	log.Printf("[Analytics] Listening at %s", s.address())
	log.Fatal(http.ListenAndServe(s.address(), h))
}

// address returns formatted hostname and port
func (s *Server) address() string {
	return fmt.Sprintf("%s:%s", s.Hostname, s.HTTPPort)
}
