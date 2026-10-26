//go:build tools

// This file never builds. waits.md is a notebook, not Go, so `go mod tidy` --
// which Renovate runs after every bump -- finds nothing here that uses the
// libraries the notebook imports and deletes their requirements. Importing
// them here keeps them, at the versions go.mod pins.
package example

import (
	_ "github.com/go-ndarray/ndarray/random"
)
