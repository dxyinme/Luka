package main

import (
	"log"
	"reflect"
	"testing"
)


func TestServer(t *testing.T){

}

func TestCall(t *testing.T){
	s:=func(a,b int)int{
		return a+b
	}
	funcValue := reflect.ValueOf(s)
	params := []reflect.Value{reflect.ValueOf(1),reflect.ValueOf(2)}
	res := funcValue.Call(params)
	log.Println(res[0].Int())
}