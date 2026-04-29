package tools

import (
	"time"
)

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"joel": {
		AuthToken: "1234567IKD",
		Username:  "joel",
	},
	"henry": {
		AuthToken: "1294AJA830",
		Username:  "henry",
	},
	"shina": {
		AuthToken: "1IFO547290",
		Username:  "shina",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"joel": {
		Coins:    100,
		Username: "joel",
	},
	"henry": {
		Coins:    100,
		Username: "henry",
	},
	"shina": {
		Coins:    100,
		Username: "shina",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
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
