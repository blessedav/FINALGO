package store

import (
	"template/internal/app/connections"
)

type ClientStore struct {
	// No clients needed for now
}

func NewClientStore(conns *connections.Connections) *ClientStore {
	return &ClientStore{}
}
