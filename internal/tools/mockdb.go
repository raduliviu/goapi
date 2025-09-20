package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"radu": {
		AuthToken: "123ABC",
		Username: "radu",
	},
	"alice": {
		AuthToken: "456DEF",
		Username: "alice",
	},
	"bob": {
		AuthToken: "789GHI",
		Username: "bob",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"radu": {
		Coins: 100,
		Username: "radu",
	},
	"alice": {
		Coins: 250,
		Username: "alice",
	},
	"bob": {
		Coins: 500,
		Username: "bob",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	clientData := LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	coinData := CoinDetails{}
	coinData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &coinData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}