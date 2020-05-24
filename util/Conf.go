package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	OK = "OK"
	FAIL = "FAIL"
)
// ServiceHost/ServicePort 是rpc提供者的
// RegisterHost/RegisterPort 是rpc发现者的
type YAML struct {
	RegisterHost string `yaml:"RegisterHost"`
	RegisterPort string `yaml:"RegisterPort"`
	RedisHost    string `yaml:"redisHost"`
	KeeperName   string `yaml:"KeeperName"`
}

var globalConf YAML

func ReadYAML(filename string) (*YAML , error) {
	yamlFile , errYAML := ioutil.ReadFile(filename)
	if errYAML != nil {
		return nil , errYAML
	}
	var conf = new(YAML)
	errYAMLParse := yaml.Unmarshal(yamlFile , conf)
	if errYAMLParse != nil {
		return nil , errYAMLParse
	}
	globalConf = *conf
	return conf,nil
}

func GetRedisHost() string {
	return globalConf.RedisHost
}