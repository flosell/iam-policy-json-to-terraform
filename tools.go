// +build tools

// Managing build tooling here. Inspired by this:
// https://marcofranssen.nl/manage-go-tools-via-go-modules/

package tools

import (
	_ "github.com/github-release/github-release"
	_ "golang.org/x/lint/golint"
)
