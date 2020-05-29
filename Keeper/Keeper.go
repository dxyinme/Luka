package Keeper

import (
	"Luka/util"
	"github.com/garyburd/redigo/redis"
	"log"
)


// keeper
type Keeper struct {
	Name          string `json:"name"`
	KeeperUrl     string `json:"keeperUrl"`
}

var redisConn redis.Conn
var set = make(map[string]bool)


func ResetRedis() {
	c,errRedis := redis.Dial("tcp", util.GetRedisHost())
	if errRedis != nil {
		log.Fatal(errRedis)
	}
	redisConn = c
	_, errRedis = redisConn.Do("FLUSHALL")
	if errRedis != nil {
		log.Fatal(errRedis)
	}
}

func SetKeeper(Name,url string) error {
	_,errRedis := redisConn.Do("SET", Name , url)
	set[Name] = true
	if errRedis != nil {
		return errRedis
	}
	return nil
}

func randOne() string {
	var c string
	for key := range set {
		c = key
		break
	}
	return c
}

func GetKeeper() (string, string){
	name := randOne()
	log.Println(name)
	url,errRedis := redis.String(redisConn.Do("GET",name))
	if errRedis != nil {
		log.Println(errRedis)
	}
	return name,url
}
