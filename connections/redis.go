package connections

import (
    "gin/config"
    "github.com/go-redis/redis"
)

var RedisClient = make(map[string]*redis.Client)

func init() {
    conf := config.LoadConf().Connections.Redis

    //base_data_cache
    redisConnection(conf.BaseDataCache, "base_data_cache")
}

func redisConnection(conf config.SectionRedisConf, confName string) {
    RedisClient[confName] = redis.NewClient(&redis.Options{
        Addr:     conf.Host,
        Password: conf.Password,
        DB:       conf.Db,
        PoolSize: conf.PoolSize,
    })
}
