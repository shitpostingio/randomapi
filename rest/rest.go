package rest

import (
	"github.com/shitpostingio/randomapi/backstore"
)

// Interface is a struct holding data needed by the REST
// interface to correctly handle its requests.
type Interface struct {
	backstore *backstore.Backstore
}

// NewInterface returns a new REST interface, instantiated with
// the backstore instance b.
func NewInterface(b *backstore.Backstore) *Interface {
	return &Interface{
		backstore: b,
	}
}
