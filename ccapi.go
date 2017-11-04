package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//URL_API main api url
const URL_API = "https://min-api.cryptocompare.com/data/"

//COINLIST url api
const COINLIST = "https://www.cryptocompare.com/api/data/coinlist/"

type coinListResp struct {
	Response     string              `json:"Response"`
	Message      string              `json:"Message"`
	BaseImageURL string              `json:"BaseImageUrl"`
	BaseLinkURL  string              `json:"BaseLinkUrl"`
	Data         map[string]CoinInfo `json:"Data"`
	Type         int                 `json:"Type"`
}

//Coins Coins map used to receive json resp from the api
type Coins map[string]CoinInfo

//CoinInfo is struct for CryptoCompare Api
type CoinInfo struct {
	ID                  string `json:"Id"`
	URL                 string `json:"Url"`
	ImageURL            string `json:"ImageUrl"`
	Name                string `json:"Name"`
	CoinName            string `json:"CoinName"`
	FullName            string `json:"FullName"`
	Algorithm           string `json:"Algorithm"`
	ProofType           string `json:"ProofType"`
	FullyPremined       string `json:"FullyPremined"`
	TotalCoinsFreeFloat string `json:"TotalCoinsFreeFloat"`
	TotalCoinSupply     string `json:"TotalCoinSupply"`
	PreMinedValue       string `json:"PreMinedValue"`
	SortOrder           string `json:"SortOrder"`
}

//PriceResult is used for fetching price data from the api
type PriceResult map[string]map[string]float64

//GetCoinList fetches data for all coins from the API
func GetCoinList() (Coins, error) {
	resp, err := makeReq(COINLIST)
	if err != nil {
		log.Println(err)
	}
	var data coinListResp
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Println(err)
	}

	return data.Data, err
}

//SyncPortfolio fetch prices from api and calculate total portfolio value
func SyncPortfolio(assets Assets, coins Coins) (portfolio Assets, err error) {
	//Check portfolio for invalid coins/entries
	portfolio = checkPortfolio(assets, coins)

	//storing coin's ticker in an array
	var a []string
	for _, asset := range portfolio {
		a = append(a, asset.Ticker)
	}

	//join into a string
	from := strings.Join(a, ",")
	to := "BTC," + Config.BaseCurrency
	url := URL_API + "pricemulti?fsyms=" + from + "&tsyms=" + to

	//get latest prices
	currentPrices, err := getCurrentPrice(url)
	if err != nil {
		log.Println("SyncPortfolio(): ", err)
		return
	}

	for i, v := range portfolio {
		portfolio[i].BTCPrice = currentPrices[v.Ticker]["BTC"]
		portfolio[i].FiatPrice = currentPrices[v.Ticker][Config.BaseCurrency]
	}

	return
}

//CheckCoins protfolio input
func checkPortfolio(assets Assets, coins Coins) (newList Assets) {
	for i, asset := range assets {
		//check if asset is in the coinlist
		if coins[asset.Ticker].Name != "" {
			//if its in the list fetch the full Name
			assets[i].Name = coins[assets[i].Ticker].CoinName
			newList = append(newList, assets[i])
		} else {
			log.Println("Asset not found in Coinlist: ", asset.Ticker)
		}
	}
	log.Println("Checking coinlist completed.")

	return
}

func getCurrentPrice(url string) (map[string]map[string]float64, error) {
	resp, err := makeReq(url)
	if err != nil {
		log.Println("getCurrentPrice(): ", err)
	}

	var data PriceResult
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Println("getCurrentPrice(): ", err)
	}

	return data, err
}

func doReq(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

func makeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := doReq(req)
	if err != nil {
		log.Println(err)
	}

	return resp, err
}
