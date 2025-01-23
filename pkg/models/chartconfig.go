package models

type ChartConfig struct {
	Ticker    string    `json:"ticker"`
	Timeframe Timeframe `json:"timeframe"`
	Range     Range     `json:"range"`
}
