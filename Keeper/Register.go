package Keeper

import (
	"Luka/util"
	"log"
	"net/http"
	"time"
)

func GetConfirm(URL string) bool {
	req, errReq := http.NewRequest("POST",
		URL , nil)
	if errReq != nil {
		log.Println(errReq)
	}
	client := http.Client{}
	resp,errHttp := client.Do(req)
	if errHttp != nil {
		log.Println(errHttp)
		return false
	}
	defer resp.Body.Close()
	body := util.ReadBody(resp.Body)
	if string(body) == util.OK {
		return true
	}else {
		return false
	}
}

func CircleConfirm(){
	go func() {
		for i := 0 ; true ; i ++ {
			log.Println(i)
			for k,v := range FuncMap {
				log.Println(k)
				isOK := GetConfirm("http://" + v.Host + v.Port + "/Confirm")
				if isOK {
					log.Println(v.Name + " " + util.OK)
				} else {
					log.Println(v.Name + " " + util.FAIL)
				}
			}
			time.Sleep(time.Second * 10)
		}
	}()
}