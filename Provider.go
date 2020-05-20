// +build ignore

package main

import (
	"Luka/Provider"
	"Luka/util"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main(){
	conf,err := util.ReadYAML("Keeper.yaml")
	if err != nil {
		log.Fatal(err)
	}
	sendTo := Provider.Server{
		Name: conf.KeeperName,
		Host: conf.ServiceHost,
		Port: conf.ServicePort,
	}
	byteSend,errJson := json.Marshal(sendTo)
	if errJson != nil {
		log.Println(errJson)
	}
	reader := bytes.NewReader(byteSend)
	req, errReq := http.NewRequest("POST",
		"http://"+conf.RegisterHost+conf.RegisterPort+"/Register", reader)
	if errReq != nil {
		log.Println(errReq)
	}
	client := http.Client{}
	resp,errHttp := client.Do(req)
	if errHttp != nil {
		log.Println(errHttp)
	}
	defer resp.Body.Close()
	body := util.ReadBody(resp.Body)
	fmt.Println(string(body))
	http.HandleFunc("/Confirm", Provider.ConfirmReply)
	_ = http.ListenAndServe(conf.ServiceHost+conf.ServicePort, nil)
}