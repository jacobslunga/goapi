package tools

import (
  "time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails {
  "alex": {
    AuthToken: "123ABC",
    Username: "alex",
  },
  "jacob": {
    AuthToken: "456EDF",
    Username: "jacob",
  },
  "joe": {
    AuthToken: "789GHI",
    Username: "joe",
  },
}

var mockCoinDetails = map[string]CoinDetails {
  "alex": {
    Coins: 100,
    Username: "alex",
  },
  "jacob": {
    Coins: 200,
    Username: "jacob",
  },
  "joe": {  
    Coins: 300,
    Username: "joe",
  },
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
  time.Sleep(1 * time.Second)

  var clientData = LoginDetails{}
  clientData, ok := mockLoginDetails[username]
  if !ok {
    return nil
  }
  return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
  time.Sleep(1 * time.Second)

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
