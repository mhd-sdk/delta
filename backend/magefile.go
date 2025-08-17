//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Dev() error {
	return sh.RunV("go", "run", "cmd/delta/delta.go")
}
func Build() error {
	return sh.RunV("go", "build", "cmd/delta/delta.go")
}

// build protobuff
func Proto() error {
	return sh.RunV("protoc", "--go_out=.", "--go-grpc_out=.", "proto/state.proto")
}
