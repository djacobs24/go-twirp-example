package internal

import (
	"context"

	"github.com/djacobs24/go-twirp-example/rpc/hats"
	"github.com/segmentio/ksuid"
	"github.com/twitchtv/twirp"
)

// Server implements the Haberdasher service
type Server struct{}

// CreateHat makes a new hat
func (s *Server) CreateHat(ctx context.Context, req *hats.CreateHatRequest) (h *hats.Hat, err error) {
	// Size must be greater than 0 inches
	if req.Size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("inches", "Hat size must be greater than 0.")
	}
	// Generate id
	id := ksuid.New().String()
	// Construct hat
	resp := &hats.Hat{
		ID:    id,
		Name:  req.Name,
		Size:  req.Size,
		Color: req.Color,
	}
	// Return hat
	return resp, nil
}

// FindHat finds a hat by id
func (s *Server) FindHat(ctx context.Context, req *hats.FindHatRequest) (h *hats.Hat, err error) {
	// Generate id
	id := req.ID
	// Construct response
	resp := &hats.Hat{
		ID:    id,
		Name:  "Name",
		Size:  &hats.Size{Inches: 22},
		Color: "Color",
	}
	// Return response
	return resp, nil
}

// DeleteHat deletes a hat by id
func (s *Server) DeleteHat(ctx context.Context, req *hats.DeleteHatRequest) (dhr *hats.DeleteHatResponse, err error) {
	// Construct response
	resp := &hats.DeleteHatResponse{}
	// Return response
	return resp, nil
}
