//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Dev() error {
	return sh.RunV("go", "run", "cmd/delta/delta.go")
}
func Build() error {
	return sh.RunV("wails", "build", "-devtools", "-o", "-tags", "webkit2_41", "delta.exe")
}
