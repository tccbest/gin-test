package connections

import (
    "gin/config"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "fmt"
    "log"
    "time"
)

var MySQLConn = make(map[string]*gorm.DB)

func init() {
    conf := config.LoadConf().Connections.MySQL

    //mioji_label
    conn(conf.MiojiLabel, "mioji_label")

    //mioji_chat_public
    conn(conf.MiojiChatPublic, "mioji_chat_public")
}

func conn(conf config.SectionMySQLConf, confName string) {
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        conf.Username,
        conf.Password,
        conf.Host,
        conf.Port,
        conf.Database,
    )

    var err error
    MySQLConn[confName], err = gorm.Open("mysql", dsn)
    if err == nil {
        MySQLConn[confName].DB().SetMaxIdleConns(10)
        MySQLConn[confName].DB().SetMaxOpenConns(10)
        MySQLConn[confName].DB().SetConnMaxLifetime(time.Minute)
        MySQLConn[confName].DB().Ping()
        MySQLConn[confName].LogMode(true)
    } else {
        log.Panic(err)
    }
}
