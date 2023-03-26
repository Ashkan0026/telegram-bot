package models

import (
	"fmt"
)

type Coin struct {
	ID             string   `json:"id"`
	Icon           string   `json:"icon"`
	Name           string   `json:"name"`
	Symbol         string   `json:"symbol"`
	Rank           int      `json:"rank"`
	Price          float64  `json:"price"`
	PriceBTC       float64  `json:"priceBtc"`
	Volume         float64  `json:"volume"`
	MarketCap      float64  `json:"marketCap"`
	AvilableSupply float64  `json:"availableSupply"`
	TotalSupply    float64  `json:"totalSupply"`
	PriceChange1h  float64  `json:"priceChange1h"`
	PriceChange1d  float64  `json:"priceChange1d"`
	PriceChange1w  float64  `json:"priceChange1w"`
	WebsiteURL     string   `json:"websiteUrl"`
	RedditURL      string   `json:"redditUrl"`
	TwitterURL     string   `json:"twitterUrl"`
	Exp            []string `json:"exp"`
}

type OneCoin struct {
	Coin *Coin `json:"coin"`
}

type Coins struct {
	Coins []*Coin `json:"coins"`
}

func (c *Coin) String() string {
	return "Name : " + c.Name +
		"\nID : " + c.ID +
		"\nPrice : " + fmt.Sprintf("%f$", c.Price) +
		"\nRank : " + fmt.Sprintf("%d", c.Rank) +
		"\nVolume : " + fmt.Sprintf("%f", c.Volume) +
		"\nMarketCap : " + fmt.Sprintf("%f", c.MarketCap) +
		"\nSymbol : " + c.Symbol +
		"\nPriceBTC : " + fmt.Sprintf("%f", c.PriceBTC) +
		"\nAvilableSupply : " + fmt.Sprintf("%f", c.AvilableSupply) +
		"\nPriceChange1h : " + fmt.Sprintf("%f", c.PriceChange1h) +
		"\nPriceChange1d : " + fmt.Sprintf("%f", c.PriceChange1d) +
		"\nPriceChange1w : " + fmt.Sprintf("%f", c.PriceChange1w) +
		"\nWebsiteURL : " + c.WebsiteURL +
		"\nRedditURL : " + c.RedditURL +
		"\nTwitterURL : " + c.TwitterURL
}
