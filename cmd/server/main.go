package main

import (
	"net/http"

	server "github.com/djacobs24/go-twirp-example/internal"
	"github.com/djacobs24/go-twirp-example/rpc/haberdasher"
)

const (
	port = ":5555"
)

func main() {

	// Make server
	server := &server.Server{}
	twirpHandler := haberdasher.NewHaberdasherServer(server, nil)

	// Start server
	http.ListenAndServe(port, twirpHandler)
}
