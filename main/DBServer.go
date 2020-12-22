package main

import (
	UserHttpRouter "github.com/dxyinme/Luka/Dao/httpRouter/User"
	"log"
	"net/http"
)

// this is the micro server for DB-Operation
// after Glamorgann-kv https://github.com/Glamorgann/Glamorgann finished
// this will be used to realize the business between kv and user.
func main() {
	UserHttpRouter.Initial()
	http.Handle("/User", UserHttpRouter.Router)

	if err := http.ListenAndServe(":12777", nil); err != nil {
		log.Fatal(err)
	}
}
