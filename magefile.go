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
	return sh.RunV("wails", "dev", "-loglevel", "error", "-tags", " webkit2_41")
}
func Build() error {
	return sh.RunV("wails", "build", "-devtools", "-o", "-tags", "webkit2_41", "delta.exe")
}

func Nsis() error {
	return sh.RunV("wails", "build", "-nsis")
}

func DeleteAppdata() error {
	dir, err := persistence.GetPersistenceDir("delta")
	if err != nil {
		return err
	}
	return os.RemoveAll(dir)
}

func ResetAppdata() error {
	p, err := persistence.New("delta")
	if err != nil {
		return err
	}
	return p.ResetPreferences()
}
