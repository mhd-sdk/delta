//go:build mage

package main

import (
	"delta/pkg/persistence"
	"os"

	"github.com/magefile/mage/sh"
)

func DevBrowser() error {
	return sh.RunV("wails", "dev", "-browser", "-loglevel", "info")
}

func Dev() error {
	return sh.RunV("wails", "dev", "-loglevel", "error")
}
func Build() error {
	return sh.RunV("wails", "build", "-devtools", "-o", "delta.exe")
}

func Nsis() error {
	return sh.RunV("wails", "build", "-nsis")
}

func ClearData() error {
	dir, err := persistence.GetPersistenceDir("delta")
	if err != nil {
		return err
	}
	return os.RemoveAll(dir)
}
