//go:build tools
// +build tools

// Managing build tooling here. Inspired by this:
// https://marcofranssen.nl/manage-go-tools-via-go-modules/

package tools

import (
	_ "github.com/mgechev/revive@v1.3.9"
	_ "github.com/securego/gosec/v2/cmd/gosec@latest"
)
