# Store
[![GoDoc](https://godoc.org/github.com/gokv/store?status.svg)](https://godoc.org/github.com/gokv/store)

Package store defines a simple interface for key-value store dependency
injection.

Not all implementations are supposed to provide every method. More importantly,
the consumers should declare their subset of required methods.

This package is not meant to be imported, but rather to document how
the methods should be implemented.

```Go
// Store defines an interface for interacting with a key-value store able to
// store JSON data in some form.
type Store interface {

	// Get retrieves a new value by key and unmarshals it to v, or returns false if
	// not found.
	// Err is non-nil if key was not found, or in case of failure.
	Get(ctx context.Context, key interface{}, v json.Unmarshaler) (ok bool, err error)

	// GetAll unmarshals to c every item in the store.
	// Err is non-nil in case of failure.
	GetAll(ctx context.Context, c Collection) error

	// Add assigns the given value to the given key if it doesn't exist already.
	// Err is non-nil if key was already present, or in case of failure.
	Add(ctx context.Context, key interface{}, v json.Marshaler) error

	// Set idempotently assigns the given value to the given key.
	// Err is non-nil in case of failure.
	Set(ctx context.Context, key interface{}, v json.Marshaler) error

	// SetWithTimeout assigns the given value to the given key, possibly
	// overwriting. The assigned key will clear after timeout. The lifespan starts
	// when this function is called.
	// Err is non-nil in case of failure.
	SetWithTimeout(ctx context.Context, key interface{}, v json.Marshaler, timeout time.Duration) error

	// SetWithDeadline assigns the given value to the given key, possibly overwriting.
	// The assigned key will clear after deadline.
	// Err is non-nil in case of failure.
	SetWithDeadline(ctx context.Context, key interface{}, v json.Marshaler, deadline time.Time) error

	// Update assigns the given value to the given key, if it exists.
	// Err is non-nil if key was not found, or in case of failure.
	Update(ctx context.Context, key interface{}, v json.Marshaler) error

	// Delete removes a key and its value from the store.
	// Err is non-nil if key was not found, or in case of failure.
	Delete(ctx context.Context, key interface{}) error

	// Ping returns a non-nil error if the Store is not healthy or if the
	// connection to the persistence is compromised.
	Ping(ctx context.Context) error

	// Close releases the resources associated with the Store.
	// Any further operation may cause panic.
	// Err is non-nil in case of failure.
	Close() error
}
```
