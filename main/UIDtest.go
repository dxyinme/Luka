package main

import (
	"github.com/dxyinme/LukaComm/util/CoHash"
	"log"
)

func main() {
	uid1 := CoHash.UID{Uid: "luka"}
	uid2 := CoHash.UID{Uid: "dxy"}
	log.Printf("%v: %v", uid1.Uid, uid1.GetHash())
	log.Printf("%v: %v", uid2.Uid, uid2.GetHash())
}
