package persistence

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

type AppData struct {
	Keys            Keys        `json:"keys"`
	Preferences     Preferences `json:"preferences"`
	Workspaces      []Workspace `json:"workspaces"`
	FavoriteTickers []string    `json:"favoriteTickers"`
}

type Workspace struct {
	Name   string `json:"name"`
	Layout []Tile `json:"layout"`
}

type Tile struct {
	ID   string
	Type TileType
	X    int
	Y    int
	W    int
	H    int
}

type TileType string

const (
	Chart TileType = "Chart"
)

type Persistence struct {
	filePath string
}
