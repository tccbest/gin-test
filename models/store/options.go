package store

import (
    "gin/connections"
)

type Options struct {
    Name  string `json:"name"`
    Value string `json:"value"`
}

func (o *Options) TableName() string {
    return "options"
}

//通过name获取
func OptionsGetByName(name string) string {
    var o Options
    noRecord := connections.MySQLClient["store"].Table(o.TableName()).Where("name = ?", name).First(&o).RecordNotFound()

    if noRecord {
        return ""
    }

    return o.Value
}

//通过names获取
func OptionsGetByNames(names []string) []*Options {
    var options []*Options
    var o *Options
    noRecord := connections.MySQLClient["store"].Table(o.TableName()).Where("name in (?)", names).Find(&options).RecordNotFound()

    if noRecord {
        return nil
    }

    return options
}
