package base_data

import (
    "gin/connections"
    "encoding/json"
)

type Exchange struct {
    Name         string  `json:"na"`
    CurrencyCode string  `json:"currency_code"`
    CoinSign     string  `json:"coin_sign"`
    Rate         float64 `json:"rate"`
    ShowName     string  `json:"show_name"`
}

func (e *Exchange) TableName() string {
    return "exchange"
}

func (e *Exchange) RedisKey() string {
    return "base_data_exchange"
}

//获取所有国家数据
func ExchangeGetAll() []*Exchange {
    exchanges := ExchangeGetAllFromRedis()

    if len(exchanges) == 0 {
        exchanges = ExchangeGetAllFromMySQL()
        exchangeSetAllToRedis(exchanges)
    }

    return exchanges
}

//从mysql中获取数据
func ExchangeGetAllFromMySQL() []*Exchange {
    e := new(Exchange)
    result := make([]*Exchange, 0)
    connections.MySQLClient["base_data"].Table(e.TableName()).Find(&result).RecordNotFound()

    return result
}

//从redis中获取数据
func ExchangeGetAllFromRedis() []*Exchange {
    e := new(Exchange)
    result := make([]*Exchange, 0)

    exchanges, err := connections.RedisClient["base_data_cache"].HGetAll(e.RedisKey()).Result()
    if err != nil {
        return result
    }

    if len(exchanges) > 0 {
        for _, exchange := range exchanges {
            var e *Exchange
            json.Unmarshal([]byte(exchange), &e)
            result = append(result, e)
        }
    }

    return result
}

//数据写入到redis中
func exchangeSetAllToRedis(exchanges []*Exchange) {
    setData := make(map[string]interface{})
    e := new(Exchange)

    for _, val := range exchanges {
        exchange, _ := json.Marshal(val)
        setData[val.CurrencyCode] = string(exchange)
    }

    connections.RedisClient["base_data_cache"].HMSet(e.RedisKey(), setData).Err()
}

//func ExchangeGetByCurrencyCode(code string) (*Exchange, error) {
//
//}
