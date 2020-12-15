package AssignUtil

import (
	"flag"
	"github.com/golang/glog"
	"io/ioutil"
	"strings"
	"sync"
)

var (
	config = flag.String("config", "", "config file")
)

type Config struct {
	mu_ sync.RWMutex
	// KV
	// IPAddress : password for worker@IPAddress
	Password map[string]string
}

var Cfg = &Config{}

func (c *Config) GetPassword(h string) (string,bool) {
	c.mu_.RLock()
	defer c.mu_.RUnlock()
	ret, ok := c.Password[h]
	return ret,ok
}

func ConfigInitial() {
	Cfg.mu_.Lock()
	defer Cfg.mu_.Unlock()
	Cfg.Password = make(map[string]string)
	text, err := ioutil.ReadFile(*config)
	if err != nil {
		glog.Fatal(err)
	}
	textLine := strings.Split(string(text), "\n")
	for _,v := range textLine {
		if len(v) <= 0 {
			continue
		}
		now := strings.Split(v," ")
		Cfg.Password[now[0]] = now[1]
		glog.Infof("Host=[%s],Password=[%s]", now[0], now[1])
	}
}

func AddNode(host, password string) {
	Cfg.mu_.Lock()
	defer Cfg.mu_.Unlock()
	Cfg.Password[host] = password
}