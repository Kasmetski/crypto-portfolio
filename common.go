package main

import (
	"github.com/kasmetski/cmcAPI"
)

//SyncAssets data with CMC Api
func SyncAssets(assets map[string]*Asset) (ua UserAssets, err error) {
	//get coins info from Coin Market Cap Api
	coins, err := cmcAPI.GetAllCoinInfo(0)
	if err != nil {
		return
	}

	//sync information for the assets
	for a := range assets {
		assets[a].priceBTC = coins[a].PriceBtc
		assets[a].priceFiat = coins[a].PriceUsd
		assets[a].priceChange24H = coins[a].PercentChange24H
		assets[a].priceChange7D = coins[a].PercentChange7D
		assets[a].sumBTC = assets[a].quantity * coins[a].PriceBtc
		assets[a].sumFiat = assets[a].quantity * coins[a].PriceUsd
	}

	return
}

//CalcPortfolio - calculate BTC and USD value of the crypto portfolio
func CalcPortfolio(m map[string]*Asset) (portfolioBTC, portfolioFIAT float64) {
	for a := range m {
		portfolioBTC += m[a].sumBTC
		portfolioFIAT += m[a].sumFiat
	}

	return
}
