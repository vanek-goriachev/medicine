//go:build tools

// Package tools contains go:generate commands for all project tools with versions stored in local go.mod file
package tools

//go:generate bash -c "GOBIN=$(dirname $(pwd))/bin go install github.com/vektra/mockery/v2@v2.52.4"
//go:generate bash -c "GOBIN=$(dirname $(pwd))/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint"
//go:generate bash -c "GOBIN=$(dirname $(pwd))/bin go install golang.org/x/tools/cmd/goimports"

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/vektra/mockery/v2/cmd"
	_ "golang.org/x/tools/cmd/goimports"
)
