package Keeper

import (
	"fmt"
)

var FuncMap = make(map[string] *Server)

func SetService(Name string , s *Server) error {
	if FuncMap[Name] != nil {
		return fmt.Errorf("this name : %s , have existed", Name)
	}
	FuncMap[Name] = s
	return nil
}