package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/djacobs24/go-twirp-example/rpc/haberdasher"
)

func main() {

	// Create protobuf client
	client := haberdasher.NewHaberdasherProtobufClient("http://localhost:5555", &http.Client{})

	// Construct request
	request := &haberdasher.CreateHatRequest{
		Name:  "Baseball Cap",
		Size:  &haberdasher.Size{Inches: 12},
		Color: "Black",
	}

	// Create hat
	hat, err := client.CreateHat(context.Background(), request)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	// Parse hat data
	id := hat.ID
	name := hat.Name
	inches := hat.Size.Inches
	color := hat.Color

	// Print hat data
	fmt.Printf("\nCreated Hat:\nID: %s\nName: %s\nSize: %d\nColor: %s\n", id, name, inches, color)
}
