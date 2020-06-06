package main

import (
	"flag"

	"github.com/halafi/go-things/go-anal/internal/server"
	"github.com/halafi/go-things/go-anal/internal/routes"
)

func main() {
	hostname := flag.String(
		"h", "0.0.0.0", "hostname",
	)
	port := flag.String(
		"p", "8080", "port",
	)
	flag.Parse()

	s := server.NewServer(*hostname, *port)
	r := routes.Routes()
	s.Run(r)
}