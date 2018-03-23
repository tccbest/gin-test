package base_data

import (
    "gin/connections"
    "errors"
    "gin/helpers"
)

type Country struct {
    Mid         int    `json:"mid"`
    Name        string `json:"name"`
    NameEn      string `json:"name_en"`
    Alias       string `json:"alias"`
    CountryCode string `json:"country_code"`
    ContinentId string `json:"continent_id"`
    HotRank     string `json:"hot_rank"`
    ShortNameCn string `json:"short_name_cn"`
    ShortNameEn string `json:"short_name_en"`
    ShowName    string `json:"show_name"`
    ShowNameEn  string `json:"show_name_en"`
    RankZh      string `json:"rank_zh"`
    RankEn      string `json:"rank_en"`
    AreaCode    string `json:"area_code"`
}

func (t *Country) TableName() string {
    return "country"
}

//获取所有国家数据
func GetAllCountries() (countries []*Country, err error) {
    var t *Country
    noRecord := connections.MySQLClient["base_data"].Table(t.TableName()).Where("status = ?", "Open").Find(&countries).RecordNotFound()

    if noRecord {
        return nil, errors.New("countries not exists")
    }

    for _, country := range countries {
        country.CountryCode = helpers.FilterNullString(country.CountryCode)
        country.ShortNameCn = helpers.FilterNullString(country.ShortNameCn)
        country.ShortNameEn = helpers.FilterNullString(country.ShortNameEn)
    }

    return countries, nil
}
