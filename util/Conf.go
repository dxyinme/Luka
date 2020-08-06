package util

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	OK   = "OK"
	FAIL = "FAIL"
)

var (
	KeeperName string
	KeeperUrl  string
	MasterUrl  string
)


type YAML struct {
	RegisterHost string `yaml:"RegisterHost"`
	RegisterPort string `yaml:"RegisterPort"`
	RedisHost    string `yaml:"redisHost"`
	RedisLife    int    `yaml:"redisLife"`
}

var globalConf YAML

func ReadYAML(filename string) (*YAML, error) {
	yamlFile, errYAML := ioutil.ReadFile(filename)
	if errYAML != nil {
		return nil, errYAML
	}
	var conf = new(YAML)
	errYAMLParse := yaml.Unmarshal(yamlFile, conf)
	if errYAMLParse != nil {
		return nil, errYAMLParse
	}
	globalConf = *conf
	return conf, nil
}

func GetRedisHost() string {
	return globalConf.RedisHost
}

func GetRedisLife() int {
	return globalConf.RedisLife
}


