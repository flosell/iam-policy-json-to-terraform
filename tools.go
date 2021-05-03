// +build tools

// Managing build tooling here. Inspired by this:
// https://marcofranssen.nl/manage-go-tools-via-go-modules/

package tools

import (
	_ "golang.org/x/lint/golint"
	_ "github.com/github-release/github-release"
)
