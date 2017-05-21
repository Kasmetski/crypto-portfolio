package main

import (
	"fmt"
)

func main() {
	myAssets := make(UserAssets)
	myAssets["bitcoin"] = &Asset{"BTC", 108, 0, 0, 0, 0, 0, 0}
	myAssets["ethereum"] = &Asset{"ETH", 108, 0, 0, 0, 0, 0, 0}
	myAssets["decred"] = &Asset{"DCR", 108, 0, 0, 0, 0, 0, 0}

	//SyncAssets with global data
	SyncAssets(myAssets)

	//calculate portfolio value
	pBTC, pFIAT := CalcPortfolio(myAssets)

	fmt.Println(pBTC, "btc")
	fmt.Println(pFIAT, "USD")
}
