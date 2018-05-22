# Store
[![GoDoc](https://godoc.org/github.com/gokv/store?status.svg)](https://godoc.org/github.com/gokv/store)

Package store defines a simple interface for key-value store dependency
injection.

Not all implementations are supposed to provide every method. More importantly,
the consumers should declare their subset of required methods.

This package is not meant to be imported, but rather to document how
the methods should be implemented.
