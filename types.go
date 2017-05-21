package main

//Asset struct
type Asset struct {
	symbol         string
	quantity       float64
	priceBTC       float64
	priceFiat      float64
	sumBTC         float64
	sumFiat        float64
	priceChange24H float64
	priceChange7D  float64
}

//UserAssets - map with all Assets
type UserAssets map[string]*Asset

//UserAssetsGlobalData - gloval information of the portfolio
type UserAssetsGlobalData struct {
	numOfAssets int
	valueBTC    float64
	valueFiat   float64
}
