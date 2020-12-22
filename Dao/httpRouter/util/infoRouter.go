package util

import "net/http"

var (
	RouterInfo = make(map[string]interface{})
)

func TestRouter(w http.ResponseWriter, r *http.Request) {
	_,_ = w.Write([]byte("Test"))
}

func ShowInfoRouter(w http.ResponseWriter, r *http.Request) {
	_,_ = w.Write(ReParseBody(RouterInfo))
}