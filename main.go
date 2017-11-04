package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Simple Console Portfolio for Digital Assets")
	fmt.Println("https://github.com/Kasmetski/crypto-portfolio")

	//read config file
	Config = ReadConfig()
	fmt.Println("config file: ", Config)
	fmt.Println("Base Currency is: ", Config.BaseCurrency)

	coins, err := GetCoinList()
	if err != nil {
		return
	}

	//portfolio := CheckCoins(Config.Assets, coins)

	portfolio, err := SyncPortfolio(Config.Assets, coins)
	if err != nil {
		log.Println("Error with portfolio sync: ", err)
		return
	}

	PrintAssets(portfolio)
}
