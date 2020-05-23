package Keeper

import (
	"fmt"
)

type Keeper struct {
	Name string
	isOnline bool
	Host string
	Port string
}

func (k *Keeper) IsOnline() bool {
	return k.isOnline
}

var userMap = make(map[string] *Keeper)

func SetKeeper(Name string , s *Keeper) error {
	if userMap[Name] != nil {
		return fmt.Errorf("this name : %s , have existed", Name)
	}
	userMap[Name] = s
	return nil
}