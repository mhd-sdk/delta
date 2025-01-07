package persistence

type UserData struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberMe"`
	System     string `json:"system"`
}

type Preferences struct {
	Theme      string      `json:"theme"`
	Workspaces []Workspace `json:"workspaces"`
}

type AppData struct {
	User        UserData    `json:"user"`
	Preferences Preferences `json:"preferences"`
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
	filePath   string
	CachedData AppData `json:"cachedData"`
}
