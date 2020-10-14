package model

import "github.com/google/uuid"

// Sample - A sample model object
type Sample struct {
	ID   uuid.UUID
	Text string
}
