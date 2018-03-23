package connections

import (
    "gin/config"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "fmt"
    "log"
    "time"
)

var MySQLClient = make(map[string]*gorm.DB)

func init() {
    conf := config.LoadConf().Connections.MySQL

    //mioji_label
    mysqlConnection(conf.MiojiLabel, "mioji_label")

    //mioji_chat_public
    mysqlConnection(conf.MiojiChatPublic, "mioji_chat_public")

    //base_data
    mysqlConnection(conf.BaseData, "base_data")

    //store
    mysqlConnection(conf.Store, "store")
}

func mysqlConnection(conf config.SectionMySQLConf, confName string) {
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        conf.Username,
        conf.Password,
        conf.Host,
        conf.Port,
        conf.Database,
    )

    var err error
    MySQLClient[confName], err = gorm.Open("mysql", dsn)
    if err == nil {
        MySQLClient[confName].DB().SetMaxIdleConns(10)
        MySQLClient[confName].DB().SetMaxOpenConns(10)
        MySQLClient[confName].DB().SetConnMaxLifetime(time.Minute)
        MySQLClient[confName].DB().Ping()
        MySQLClient[confName].LogMode(true)
    } else {
        log.Panic(err)
    }
}
