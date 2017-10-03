package main

import (
	"fmt"
)

func main() {
	fmt.Println("Simple Console Portfolio for Digital Assets")

	//read config file
	Config = ReadConfig()
	fmt.Println("config file: ", Config)
	fmt.Println("Base Currency is: ", Config.BaseCurrency)

	coins, err := GetCoinList()
	if err != nil {
		return
	}

	portfolio := CheckCoins(Config.Assets, coins)

	portfolio, err = SyncPortfolio(portfolio)
	if err != nil {
		return
	}

	PrintAssets(portfolio)
}
