//

package api

import (
    "github.com/gin-gonic/gin"
    "gin/models/base-data"
)

type RespApiTest struct {
    Currency []*Currency `json:"ccy"`
}

type Currency struct {
    Code string  `json:"code"`
    Sign string  `json:"sign"`
    Rate float64 `json:"rate"`
    Name string  `json:"name"`
}

func ApiTest(c *gin.Context) gin.H {
    exchanges := base_data.ExchangeGetAll()
    ccy := map[string]string{
        "CNY": "",
        "USD": "",
        "EUR": "",
        "GBP": "",
        "AUD": "",
        "CAD": "",
        "HKD": "",
        "JPY": "",
        "KRW": "",
        "CHF": "",
        "NZD": "",
    }

    var resp RespApiTest
    for _, val := range exchanges {
        if _, ok := ccy[val.CurrencyCode]; ok {
            currency := &Currency{
                Code: val.CurrencyCode,
                Name: val.ShowName,
                Sign: val.CoinSign,
                Rate: val.Rate,
            }
            resp.Currency = append(resp.Currency, currency)
        }
    }

    return gin.H{
        "c": 0,
        "m": "",
        "d": resp,
    }
}
