package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Ashkan0026/telegram-bot1/models"
)

func GetCurrencies() []*models.Coin {
	coins := &models.Coins{}
	res, err := http.Get("https://api.coinstats.app/public/v1/coins?skip=0&limit=20")
	if err != nil {
		log.Println("Error while fetching data")
	}
	ConvertToCoin(coins, res)
	return coins.Coins
}

func ConvertToCoin(x interface{}, res *http.Response) {
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("can not read the body")
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		log.Printf("Error while unmarshaling the data %v\n", err)
	}
}

func GetCurrency(name string) *models.OneCoin {
	name = strings.ToLower(name)
	coin := &models.OneCoin{}
	res, err := http.Get("https://api.coinstats.app/public/v1/coins/" + name)
	if err != nil {
		log.Println("Error while fetching data")
		return nil
	}
	ConvertToCoin(coin, res)
	return coin
}
