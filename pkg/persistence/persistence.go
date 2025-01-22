package persistence

import (
	"delta/pkg/models"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Persistence struct {
	filePath string
}

// Get standard config directory for the current OS
func GetPersistenceDir(appName string) (string, error) {
	var dir string

	// Detect OS
	configDir, err := os.UserConfigDir()
	if err == nil {
		// On Windows, returns %models.APPDATA%\<appName>
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

func (p *Persistence) Save(data models.AppData) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p.filePath, file, 0644)
}

func (p *Persistence) Load() (models.AppData, error) {
	var data models.AppData
	file, err := os.ReadFile(p.filePath)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(file, &data)
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

	DefaultData := models.AppData{
		Preferences: models.Preferences{
			GeneralPreferences: models.GeneralPreferences{
				Theme: models.LightTheme,
			},
		},
	}

	_, err = p.Load()
	// If file not found, create it
	if err != nil {
		err = p.Save(DefaultData)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (p *Persistence) ResetPreferences() error {
	old, err := p.Load()
	if err != nil {
		return err
	}
	DefaultData := models.AppData{
		Keys: models.Keys{
			ApiKey:    old.Keys.ApiKey,
			SecretKey: old.Keys.SecretKey,
		},
		Preferences: models.Preferences{
			GeneralPreferences: models.GeneralPreferences{
				Theme: models.LightTheme,
			},
		},
		Workspaces:      []models.Workspace{},
		FavoriteTickers: []string{},
	}
	return p.Save(DefaultData)
}
