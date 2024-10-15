//go:build mage

package main

import (
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
)

func Run() error {
	// run and keep logs
	return sh.RunV("go", "run", "main.go")
}

func Build() error {
	return sh.Run("go", "build", "-o", "server", "server.go")
}

func Proto() error {
	// Chemin vers le dossier contenant les .proto
	protoDir := "rithmic/proto/"
	// Dossier de sortie pour les fichiers générés
	outputDir := "./pkg/generated"

	// Créer une slice pour stocker les fichiers .proto trouvés
	var protoFiles []string

	// Parcours du dossier pkg/proto pour trouver les fichiers .proto
	err := filepath.Walk(protoDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Si le fichier a l'extension .proto, l'ajouter à la slice
		if filepath.Ext(path) == ".proto" {
			protoFiles = append(protoFiles, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Exécuter protoc pour chaque fichier trouvé
	for _, file := range protoFiles {
		err := sh.Run("protoc", "--go_out="+outputDir, "--go-grpc_out="+outputDir, "--proto_path="+protoDir, file)
		if err != nil {
			return err
		}
	}

	return nil
}
