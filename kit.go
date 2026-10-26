//go:build tools

// Package site keeps the landing kit in go.mod. A Hugo site has no Go code,
// so without this import `go mod tidy` (which Renovate runs after every bump)
// would find nothing that uses the module and delete its requirement. The
// build tag keeps the file out of every build.
package site

import _ "github.com/go-fleettools/landingkit"
