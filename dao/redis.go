package dao

import (
	"github.com/go-redis/redis"
)

var RD *redis.Client

func RedisInit()(err error){
	var redisdb *redis.Client
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}
	RD=redisdb
	return nil
}
func RedisClose() {
	RD.Close()
}
