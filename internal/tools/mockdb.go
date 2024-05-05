package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"vinicius": {
		AuthToken: "123ABC",
		Username:  "vinicius",
	},
	"jimmy": {
		AuthToken: "456DEF",
		Username:  "jimmy",
	},
	"chorrito": {
		AuthToken: "789GHI",
		Username:  "chorrito",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"vinicius": {
		Coins:    5000,
		Username: "vinicius",
	},
	"jimmy": {
		Coins:    4000,
		Username: "jimmy",
	},
	"chorito": {
		Coins:    603,
		Username: "chorrito",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}

	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
