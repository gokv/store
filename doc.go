/*
Package store strives to define a portable interface for key-value store
dependency injection.

Not all implementations are supposed to provide every method.

The consumer is advised to define an interface with its own subset of required
methods.

This package mainly exists to document how the methods are supposed to be
implemented.
*/
package store // import "github.com/gokv/store"

import (
	"context"
	"encoding/json"
	"time"
)

// Store defines an interface for interacting with a key-value store able to
// store JSON data in some form.
type Store interface {

	// Get retrieves a new value by key and unmarshals it to v.
	// Ok is false if the key was not found.
	// Err is non-nil in case of failure.
	Get(ctx context.Context, k string, v json.Unmarshaler) (ok bool, err error)

	// GetAll unmarshals to c every item in the store.
	// Err is non-nil in case of failure.
	GetAll(ctx context.Context, c Collection) error

	// Add assigns the given value to a new key, and returns the key.
	// Err is non-nil in case of failure.
	Add(ctx context.Context, v json.Marshaler) (k string, err error)

	// Set idempotently assigns the given value to the given key.
	// Err is non-nil in case of failure.
	Set(ctx context.Context, k string, v json.Marshaler) error

	// SetWithTimeout assigns the given value to the given key, possibly
	// overwriting. The assigned key will clear after timeout. The lifespan starts
	// when this function is called.
	// Err is non-nil in case of failure.
	SetWithTimeout(ctx context.Context, k string, v json.Marshaler, timeout time.Duration) error

	// SetWithDeadline assigns the given value to the given key, possibly overwriting.
	// The assigned key will clear after deadline.
	// Err is non-nil in case of failure.
	SetWithDeadline(ctx context.Context, k string, v json.Marshaler, deadline time.Time) error

	// Update assigns the given value to the given key, if it exists.
	// Ok is false if the key was not found.
	// Err is non-nil in case of failure.
	Update(ctx context.Context, k string, v json.Marshaler) (ok bool, err error)

	// Delete removes a key and its value from the store.
	// Ok is false if the key was not found.
	// Err is non-nil in case of failure.
	Delete(ctx context.Context, k string) (ok bool, err error)

	// Ping returns a non-nil error if the Store is not healthy or if the
	// connection to the persistence is compromised.
	Ping(ctx context.Context) error

	// Close releases the resources associated with the Store.
	// Any further operation may cause panic.
	// Err is non-nil in case of failure.
	Close() error
}

// New returns an unmarshaler for the Store to unmarshal the next fetched item into.
// New allows a collection type (e.g. a slice) to be used as an argument
// to a query (e.g. store.GetAll) to collect multiple results.
type Collection interface {
	New() json.Unmarshaler
}
