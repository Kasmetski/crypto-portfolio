package main

import (
	"fmt"
	"strings"
)

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

	fmt.Println(strings.Repeat("-", 53))
	fmt.Printf("|%11s|%6s|%11s|%8s|%11s|\n", "Name", "Ticker", "Quantity", "BTCPrice", Config.BaseCurrency)

	for _, coin := range assets {
		fmt.Printf("|%11s|%6s|%11.03f|%8f|%11.2f|\n", coin.Name, coin.Ticker, coin.Qty, coin.BTCPrice, coin.FiatPrice)

		totalBTCValue += coin.Qty * coin.BTCPrice
		totalFiatValue += coin.Qty * coin.FiatPrice
	}
	fmt.Println(strings.Repeat("-", 53))
	fmt.Printf("Total value : %.8f BTC\n", totalBTCValue)
	fmt.Printf("Total value : %.2f %s\n", totalFiatValue, Config.BaseCurrency)
	fmt.Println(strings.Repeat("-", 53))
}
