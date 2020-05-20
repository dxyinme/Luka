package Keeper

import (
	"Luka/util"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Name string `json:"name"`
}

func Register(writer http.ResponseWriter, request *http.Request) {
	errParse := request.ParseForm()
	if errParse != nil {
		log.Println(errParse)
	}
	s,errRead := ioutil.ReadAll(request.Body)
	if errRead != nil {
		log.Println(errRead)
	}
	log.Println(string(s))
	var newService Server
	errJson := json.Unmarshal(s, &newService)
	if errJson != nil {
		log.Println(errJson)
	}
	errSetService := SetService(newService.Name, &newService)
	if errSetService != nil {
		log.Println(errSetService)
	}
	_, _ = writer.Write([]byte(util.OK))
}