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
    MySQL SectionMySQL `yaml:"mysql"`
    Redis SectionRedis `yaml:"redis"`
    Mongodb SectionMongodb `yaml:"mongodb"`
}

type SectionMySQL struct {
    Test SectionTest `yaml:"test"`
}

type SectionRedis struct {
    Test SectionTest `yaml:"test"`
}

type SectionMongodb struct {
    Test SectionTest `yaml:"test"`
}

type SectionTest struct {
    Host     string `yaml:"host"`
    Port     string `yaml:"port"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
    Database string `yaml:"database"`
}

func LoadConf() {
    m := multiconfig.NewWithPath("config/config.yml")

    serverConf := new(Config)
    err := m.Load(serverConf)
    if err != nil {
        panic(err)
    }

    m.MustLoad(serverConf)
}
