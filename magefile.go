//go:build mage

package main

import (
	"delta/pkg/persistence"
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
)

func Dev() error {
	return sh.RunV("wails", "dev", "-browser")
}

func ClearData() error {
	dir, err := persistence.GetPersistenceDir("delta")
	if err != nil {
		return err
	}
	return os.RemoveAll(dir)
}

func Proto() error {

	protoDir := "docs/rithmic/proto/"
	outputDir := "./pkg/generated"

	var protoFiles []string

	err := filepath.Walk(protoDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".proto" {
			protoFiles = append(protoFiles, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	for _, file := range protoFiles {
		err := sh.Run("protoc", "--go_out="+outputDir, "--go-grpc_out="+outputDir, "--proto_path="+protoDir, file)
		if err != nil {
			return err
		}
	}

	return nil
}
