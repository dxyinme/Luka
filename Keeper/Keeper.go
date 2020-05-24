package Keeper

import (
	"Luka/util"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

type Keeper struct {
	Name     string `json:"name"`
	IsOnline bool   `json:"isOnline"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func (k *Keeper) checkOnline() bool {
	return k.IsOnline
}

var redisConn redis.Conn

func ResetRedis() {
	c,errRedis := redis.Dial("tcp", util.GetRedisHost())
	if errRedis != nil {
		log.Fatal(errRedis)
	}
	redisConn = c
}

func SetKeeper(Name string , s *Keeper) error {
	_,errCK := redis.String(redisConn.Do("GET",Name))
	if errCK != redis.ErrNil {
		return fmt.Errorf("this %s has existed" , Name)
	}
	byteStr,errJson := json.Marshal(s)
	if errJson != nil {
		return errJson
	}
	str := string(byteStr)
	_,errRedis := redisConn.Do("SET",Name,str)
	if errRedis != nil {
		return errRedis
	}
	return nil
}

