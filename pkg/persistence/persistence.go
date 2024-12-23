package persistence

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Get standard config directory for the current OS
func GetPersistenceDir(appName string) (string, error) {
	var dir string

	// Detect OS
	configDir, err := os.UserConfigDir()
	if err == nil {
		// On Windows, returns %APPDATA%\<appName>
		// On Unix, returns ~/.config/<appName>
		dir = filepath.Join(configDir, appName)
	} else {
		// Fallback to current directory in case of error
		dir = "."
	}
	// Create directory if it doesn't exist
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	return dir, nil
}

func (p *Persistence) Save(data AppData) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	p.CachedData = data
	return os.WriteFile(p.filePath, file, 0644)
}

func (p *Persistence) Load() (AppData, error) {
	var data AppData
	file, err := os.ReadFile(p.filePath)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(file, &data)
	p.CachedData = data
	return data, err
}

func New(appName string) (*Persistence, error) {
	dir, err := GetPersistenceDir(appName)
	if err != nil {
		fmt.Println("Error getting persistence directory:", err)
		return nil, err
	}

	saveFile := filepath.Join(dir, appName+".json")

	p := &Persistence{
		filePath: saveFile,
	}

	DefaultData := AppData{}

	_, err = p.Load()
	// If file not found, create it
	if err != nil {
		err = p.Save(DefaultData)
	}

	return p, nil
}
