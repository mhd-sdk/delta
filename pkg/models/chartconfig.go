package models

type ChartConfig struct {
	Ticker    string    `json:"ticker"`
	Timeframe TimeFrame `json:"timeframe"`
	Range     Range     `json:"range"`
}
