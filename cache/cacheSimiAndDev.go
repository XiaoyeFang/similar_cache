package cache

import (
	"fmt"
	"similar_cache/config"
	"similar_cache/models"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
)

func HandlerCache(key string, value interface{}) (err error) {
	//只缓存 similar_developer 接口
	//SSDB  缓存10h，当缓存中有数据時，不再去查询mongodb
	//key的起名加前缀
	conn, err := models.ConnRedis()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	_, err = conn.Do("SET", key, value, "EX", config.CacheConfig.RedisDB.Expiration)
	if err != nil {
		glog.V(0).Infoln(err)
	}

	return err
}

func Getcache(key string) (val []byte, err error) {

	conn, err := models.ConnRedis()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	val, err = redis.Bytes(conn.Do("GET", key))

	if err != nil {
		glog.V(0).Infoln(err)
		return val, err
	}
	if len(val) == 0 {
		return val, err
	}

	return val, err
}

func DeleteRediskey() {
	client, err := models.ConnRedis()
	if err != nil {
		panic(err)
	}
	_, err = client.Do("DEL", "similar_developer_v2com.netease.mobimail")
	if err != nil {
		fmt.Println("DEL err", err)
	}
}
