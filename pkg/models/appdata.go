package models

type AppData struct {
	Keys            Keys        `json:"keys"`
	Preferences     Preferences `json:"preferences"`
	Workspaces      []Workspace `json:"workspaces"`
	FavoriteTickers []string    `json:"favoriteTickers"`
}

type Preferences struct {
	GeneralPreferences GeneralPreferences `json:"generalPreferences"`
}

type Theme string

const (
	LightTheme Theme = "light"
	DarkTheme  Theme = "dark"
)

type Language string

const (
	EnglishLanguage Language = "en"
	FrenchLanguage  Language = "fr"
)

type GeneralPreferences struct {
	Theme    Theme    `json:"theme"`
	Language Language `json:"language"`
}

type Keys struct {
	ApiKey    string `json:"apiKey"`
	SecretKey string `json:"secretKey"`
}
