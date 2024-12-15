package persistence

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Preferences struct {
	Theme string `json:"theme"`
}

type AppData struct {
	User        UserData    `json:"user"`
	Preferences Preferences `json:"preferences"`
}

type Persistence struct {
	filePath   string
	CachedData AppData `json:"cachedData"`
}
