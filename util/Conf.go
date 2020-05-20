package util

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
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
	ServiceHost string `yaml:"ServiceHost"`
	ServicePort string `yaml:"ServicePort"`
	KeeperName string `yaml:"KeeperName"`
}

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
	return conf,nil
}

// 读取body
func ReadBody(bodyIo io.Reader) []byte {
	body,errRESP := ioutil.ReadAll(bodyIo)
	if errRESP != nil {
		log.Println(errRESP)
	}
	return body
}