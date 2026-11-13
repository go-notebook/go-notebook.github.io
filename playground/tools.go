//go:build tools

// Package playground pins the version of github.com/go-notebook/playground
// that the deploy workflow builds into /playground/: go.mod and go.sum say
// which, and Renovate moves it. The notebooks it shows are in notebooks/;
// the build runs each one first and fails if a cell fails.
package playground

import (
	_ "github.com/go-notebook/playground/cmd/build"
	// The build compiles cmd/editor (js/wasm, since v0.11.0) in this module,
	// so its imports must be in go.sum too.
	_ "github.com/go-notebook/playground/cmd/editor"
)
