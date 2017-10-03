package main

import "fmt"

//Asset struct
type Asset struct {
	Name      string
	Ticker    string
	Qty       float64
	BTCPrice  float64 //current BTC price
	FiatPrice float64 //current fiat price
}

//Assets array struct
type Assets []Asset

//PrintAssets used for simple portfolio view
func PrintAssets(assets Assets) {
	var totalBTCValue, totalFiatValue float64
	fmt.Println("--------")
	fmt.Printf("|%11s|%6s|%11s|%8s|%8s|\n", "Name", "Ticker", "Quantity", "BTCPrice", Config.BaseCurrency)
	for _, coin := range assets {
		fmt.Printf("|%11s|%6s|%11f|%8f|%8.2f|\n", coin.Name, coin.Ticker, coin.Qty, coin.BTCPrice, coin.FiatPrice)

		totalBTCValue += coin.Qty * coin.BTCPrice
		totalFiatValue += coin.Qty * coin.FiatPrice
	}
	fmt.Println("--------")
	fmt.Println("Total value (BTC): ", totalBTCValue)
	fmt.Println("Total value (FIAT): ", totalFiatValue, Config.BaseCurrency)
	fmt.Println("--------")
}
