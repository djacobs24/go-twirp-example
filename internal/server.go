package internal

import (
	"context"
	"time"

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
	// Generate id and timestamps
	id := ksuid.New().String()
	createdAt := time.Now().Unix()
	updatedAt := createdAt
	// Construct hat
	resp := &hats.Hat{
		ID:        id,
		Name:      req.Name,
		Size:      req.Size,
		Color:     req.Color,
		Active:    true,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	// Return hat
	return resp, nil
}

// FindHat finds a hat by id
func (s *Server) FindHat(ctx context.Context, req *hats.FindHatRequest) (h *hats.Hat, err error) {
	// Generate id and timestamps
	id := req.ID
	size := &hats.Size{Inches: 22}
	createdAt := time.Now().Unix()
	updatedAt := createdAt
	// Construct response
	resp := &hats.Hat{
		ID:        id,
		Name:      "Name",
		Size:      size,
		Color:     "Color",
		Active:    true,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
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
