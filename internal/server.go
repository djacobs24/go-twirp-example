package internal

import (
	"context"

	"github.com/djacobs24/go-twirp-example/rpc/haberdasher"
	"github.com/segmentio/ksuid"
	"github.com/twitchtv/twirp"
)

// Server implements the Haberdasher service
type Server struct{}

// CreateHat makes a new hat
func (s *Server) CreateHat(ctx context.Context, chr *haberdasher.CreateHatRequest) (h *haberdasher.Hat, err error) {

	// Size must be greater than 0 inches
	if chr.Size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("inches", "Hat size must be greater than 0.")
	}

	// Generate id
	id := ksuid.New().String()

	// Construct hat
	ht := &haberdasher.Hat{
		ID:    id,
		Name:  chr.Name,
		Size:  chr.Size,
		Color: chr.Color,
	}

	// Return hat
	return ht, nil
}
