package main

import (
	"Luka/util"
	"log"
)


func main(){
	conf,err := util.ReadYAML("Register.yaml")
	if err != nil {
		log.Println(err)
	}
	log.Println(conf)
	log.Println("hello Register!!!")
}