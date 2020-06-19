// +build ignore

package main

import (
	"Luka/Keeper"
	"log"
	"net/http"
)
// 一个 Keeper 有且只能有一个 Connector
func main(){
	newKeeper := Keeper.NewConnector(
		// 跨域
		func(r *http.Request) bool {
			return true
		})
	http.HandleFunc("/ConnectIt",newKeeper.ConnectIt)
	if err := http.ListenAndServe(":10137", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}