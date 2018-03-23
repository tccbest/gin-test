package config

import (
    "github.com/koding/multiconfig"
)

type Config struct {
    Connections SectionConnections `yaml:"connections"`
    Params      SectionParams      `yaml:"params"`
}

type SectionParams struct {
    TestParams string `yaml:"test_params"`
}

type SectionConnections struct {
    MySQL   SectionMySQL   `yaml:"mysql"`
    Redis   SectionRedis   `yaml:"redis"`
    Mongodb SectionMongodb `yaml:"mongodb"`
}

type SectionMySQL struct {
    MiojiLabel      SectionMySQLConf `yaml:"mioji_label"`
    MiojiChatPublic SectionMySQLConf `yaml:"mioji_chat_public"`
    BaseData        SectionMySQLConf `yaml:"base_data"`
    Store           SectionMySQLConf `yaml:"store"`
}

type SectionRedis struct {
    Test          SectionTest      `yaml:"test"`
    BaseDataCache SectionRedisConf `yaml:"base_data_cache"`
}

type SectionMongodb struct {
    Test1 SectionTest1 `yaml:"test"`
}

type SectionMySQLConf struct {
    Host     string `yaml:"host"`
    Port     string `yaml:"port"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
    Database string `yaml:"database"`
}

type SectionRedisConf struct {
    Host     string `yaml:"host"`
    Password string `yaml:"password"`
    Db       int    `yaml:"db"`
    PoolSize int    `yaml:"pool_size"`
}

type SectionTest struct {
    Host     string `yaml:"host"`
    Port     string `yaml:"port"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
    Database string `yaml:"database"`
}

type SectionTest1 struct {
    Uri string `yaml:"uri"`
}

func LoadConf() *Config {
    m := multiconfig.NewWithPath("config/config.yml")

    serverConf := new(Config)
    err := m.Load(serverConf)
    if err != nil {
        panic(err)
    }

    m.MustLoad(serverConf)

    return serverConf
}
