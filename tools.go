//go:build tools
// +build tools

package tools

import (
	_ "github.com/fjl/gencodec" // stop removing this from go.mod we need it for go:gen, this line can't me in main because then it gets removed because it is a 'tool' because it declares a main package. check out go 1.24 release notes for tools.go special meaning
)
