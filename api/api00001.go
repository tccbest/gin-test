//

package api

import (
    "github.com/gin-gonic/gin"
    "gin/models/store"
    "gin/helpers"
    "encoding/json"
)

type RespApi00001 struct {
    Continent  []*Continent  `json:"continent"`
    HotelBrand []*HotelBrand `json:"hotel_brand"`
    HotelScore []*HotelScore `json:"hotel_score"`
}

type Continent struct {
    Id     string `json:"id"`
    Name   string `json:"name"`
    NameEn string `json:"name_en"`
}

type HotelBrand struct {
    Id   string `json:"id"`
    Name string `json:"name"`
}

type HotelScore struct {
    Sid   string  `json:"sid"`
    Name  string  `json:"name"`
    Img   string  `json:"img"`
    Score float64 `json:"score"`
    Desc  string  `json:"desc"`
}

func Api00001(c *gin.Context) gin.H {
    options := store.OptionsGetByNames([]string{"hotel_brand", "hotel_score"})
    var hotelBrand []*HotelBrand
    var hotelScore []*HotelScore

    if options != nil {
        for _, v := range options {
            if v.Name == "hotel_brand" {
                json.Unmarshal([]byte(v.Value), &hotelBrand)
            } else if v.Name == "hotel_score" {
                json.Unmarshal([]byte(v.Value), &hotelScore)
            }
        }
    }

    Continent := []*Continent{
        {"10", "亚洲", "Asia"},
        {"20", "欧洲", "Europe"},
        {"50", "北美洲", "North America"},
        {"60", "南美洲", "South America"},
        {"40", "非洲", "Africa"},
        {"30", "大洋洲", "Oceania"},
    }

    return helpers.Response(RespApi00001{
        Continent, hotelBrand, hotelScore,
    })
}
