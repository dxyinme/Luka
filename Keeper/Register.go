package Keeper

import (
	"log"
	"time"
)

func CircleConfirm(){
	for i := 0 ; true ; i ++ {
		log.Println(i)

		time.Sleep(time.Second * 10)
	}
}