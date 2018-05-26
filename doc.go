/*
Package store defines a simple interface for key-value store dependency
injection.

Not all implementations are supposed to provide every method. More importantly,
the consumers should declare their subset of required methods.

This package is not meant to be imported, but rather to document how
the methods should be implemented.
*/
package store // import "github.com/gokv/store"

import (
	"context"
	"encoding"
	"time"
)

// Store defines a key-value store interface.
// All methods should be implemented concurrent-safe.
type Store interface {

	// Get retrieves a new object by key and unmarshals it into v, or returns false
	// if not found.
	// Err is non-nil in case of failure.
	Get(key string, v encoding.BinaryUnmarshaler) (ok bool, err error)

	// Set persists a new object, possibly overwriting.
	// Err is non-nil in case of failure.
	Set(key string, v encoding.BinaryMarshaler) error

	// Add persists a new object.
	// Err is non-nil if key is already present, or in case of failure.
	Add(key string, v encoding.BinaryMarshaler) error

	// SetWithTimeout assigns the given value to the given key, possibly
	// overwriting. The assigned key will clear after timeout. The lifespan starts
	// when this function is called.
	// Err is non-nil in case of failure.
	SetWithTimeout(key string, v encoding.BinaryMarshaler, timeout time.Duration) error

	// SetWithDeadline assigns the given value to the given key, possibly overwriting.
	// The assigned key will clear after deadline.
	// Err is non-nil in case of failure.
	SetWithDeadline(key string, v encoding.BinaryMarshaler, deadline time.Time) error

	// Keys lists all the stored keys.
	Keys(context.Context) (keys <-chan string, errs <-chan error)

	// Del idempotently removes a key from the store.
	// Err is non-nil in case of failure.
	Del(key string) error

	// Ping returns a non-nil error if the store is not healthy.
	Ping() error

	// Close releases the resources associated with the store.
	// Any further operation may cause panic.
	// Err is non-nil in case of failure.
	Close() error
}
