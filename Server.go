package main

import (
	"Luka/Keeper"
	"Luka/util"
	"log"
	"net/http"
)


func main(){
	conf,err := util.ReadYAML("Register.yaml")
	if err != nil {
		log.Println(err)
	}
	log.Println(conf)
	log.Println("hello Register!!!")
	Keeper.CircleConfirm()
	http.HandleFunc("/Register", Keeper.Register)
	_ = http.ListenAndServe(conf.RegisterHost+conf.RegisterPort, nil)
}