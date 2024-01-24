//go:build tools
// +build tools

package main

import (
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/zitadel/oidc/v3/example/server"
)
