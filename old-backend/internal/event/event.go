package event

import "time"

type Event struct {
	Alerts []Alert
	Time   time.Time
	Symbol string
}

type Alert int

const (
	UnusualVolume Alert = iota
	HighVolatility
)

type Events []Event

func (events *Events) IsUnique(symbol string) bool {
	for _, event := range *events {
		if event.Symbol == symbol {
			return true
		}
	}
	return false
}
