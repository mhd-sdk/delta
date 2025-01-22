package models

type Tile struct {
	ID   string   `json:"id"`
	X    int      `json:"x"`
	Y    int      `json:"y"`
	W    int      `json:"w"`
	H    int      `json:"h"`
	Data TileData `json:"data"`
}

type TileData struct {
	Type TileType    `json:"type"`
	Data interface{} `json:"data"`
}

type TileType string

const (
	Chart           TileType = "Chart"
	TimeAndSales    TileType = "TimeAndSales"
	AccountOverview TileType = "AccountOverview"
	TickerInfo      TileType = "TickerInfo"
	Headlines       TileType = "Headlines"
	Scanner         TileType = "Scanner"
)
