package srv

import "delta/pkg/services/market"

type Srv struct {
	services struct {
		marketService market.MarketService
	}
}
