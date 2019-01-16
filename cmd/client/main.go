package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/djacobs24/go-twirp-example/rpc/hats"
	"github.com/segmentio/ksuid"
)

func main() {
	// Create protobuf client
	client := hats.NewHatsProtobufClient("http://localhost:5555", &http.Client{})
	createHat(client)
	findHat(client)
	deleteHat(client)
}

func createHat(client hats.Hats) {
	// Construct request
	request := &hats.CreateHatRequest{
		Name:  "Baseball Cap",
		Size:  &hats.Size{Inches: 22},
		Color: "Black",
	}
	// Create hat
	hat, err := client.CreateHat(context.Background(), request)
	if err != nil {
		fmt.Printf("Create Hat Error: %v", err)
		return
	}
	// Parse hat data
	id := hat.ID
	name := hat.Name
	inches := hat.Size.Inches
	color := hat.Color
	active := hat.Active
	createdAt := hat.CreatedAt
	updatedAt := hat.UpdatedAt
	// Print hat data
	fmt.Printf("\nCreated Hat:\nID: %s\nName: %s\nSize: %d\nColor: %s\nActive: %t\nCreated At: %d\nUpdated At: %d\n", id, name, inches, color, active, createdAt, updatedAt)
}

func findHat(client hats.Hats) {
	// Generate id
	id := ksuid.New().String()
	// Construct request
	request := &hats.FindHatRequest{
		ID: id,
	}
	// Find hat
	resp, err := client.FindHat(context.Background(), request)
	if err != nil {
		fmt.Printf("Find Hat Error: %v", err)
		return
	}
	// Parse hat data
	name := resp.Name
	inches := resp.Size.Inches
	color := resp.Color
	active := resp.Active
	createdAt := resp.CreatedAt
	updatedAt := resp.UpdatedAt
	// Print hat data
	fmt.Printf("\nFound Hat:\nID: %s\nName: %s\nSize: %d\nColor: %s\nActive: %t\nCreated At: %d\nUpdated At: %d\n", id, name, inches, color, active, createdAt, updatedAt)
}

func deleteHat(client hats.Hats) {
	// Generate id
	id := ksuid.New().String()
	// Construct request
	request := &hats.DeleteHatRequest{
		ID: id,
	}
	// Delete hat
	_, err := client.DeleteHat(context.Background(), request)
	if err != nil {
		fmt.Printf("Delete Hat Error: %v", err)
		return
	}
	// Print hat data
	fmt.Printf("\nDeleted Hat:\nID: %s\n", id)
}
