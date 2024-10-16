package market

import (
	"delta/pkg/generated/rti"
	"delta/pkg/services"
)

type MarketServiceImpl interface {
}

type MarketService struct {
	services.RithmicWS
}

func NewMarketService() *MarketService {
	mktService := MarketService{}
	args := services.ConnectionArgs{}
	err := mktService.Connect(args, rti.RequestLogin_TICKER_PLANT)
	if err != nil {
		panic(err)
	}
	return &mktService
}
