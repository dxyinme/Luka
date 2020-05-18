package util

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func transform(param string, typeName string) (interface{},error) {
	switch typeName {
	case "string": {
		return param,nil
	}
	case "int":{
		now,err := strconv.Atoi(param)
		if err != nil {
			return nil,err
		}
		return now,nil
	}
	case "int32":{
		now,err := strconv.ParseInt(param,10,32)
		if err != nil {
			return nil,err
		}
		return now, nil
	}
	case "int64":{
		now,err := strconv.ParseInt(param,10,64)
		if err != nil {
			return nil,err
		}
		return now, nil
	}
	case "float32":{
		now,err := strconv.ParseFloat(param,32)
		if err != nil {
			return nil,err
		}
		return now,nil
	}
	case "float64":{
		now,err := strconv.ParseFloat(param,64)
		if err != nil {
			return nil,err
		}
		return now, nil
	}
	case "bool":{
		now,err := strconv.ParseBool(param)
		if err != nil {
			return nil,err
		}
		return now, nil
	}
	default:{
		return nil, fmt.Errorf("there is no such type %s", typeName)
	}
	}

}

// 将字符串数组转化成reflect参数数组
func TransformList(paramList []string, typeList []string) ([]interface{},error) {
	length := len(paramList)
	if length != len(typeList) {
		return nil , fmt.Errorf("error Length")
	}
	res := make([]interface{}, length)
	var err error = nil
	for i := 0 ; i < length ; i++ {
		res[i],err = transform(paramList[i] , typeList[i])
		if err != nil {
			return nil,err
		}
	}
	return res,nil
}

// 将接口类型转换成reflect.Value类型
func TransIntoReflect(v []interface{}) []reflect.Value {
	length := len(v)
	ret := make([]reflect.Value,length)
	for  i := 0 ; i < length ; i++  {
		ret[i] = reflect.ValueOf(v[i])
	}
	return ret
}


func transReflectIntoString(v reflect.Value) (string,string) {
	switch v.Type().Name() {
	case "string":{
		return v.String(),"string"
	}
	case "int":{
		return strconv.Itoa(int(v.Int())),"int"
	}
	case "int32":{
		return strconv.FormatInt(v.Int(),10),"int32"
	}
	case "int64":{
		return strconv.FormatInt(v.Int(),10),"int64"
	}
	case "float32":{
		return strconv.FormatFloat(v.Float(),'E',-1,32),"float32"
	}
	case "float64":{
		return strconv.FormatFloat(v.Float(),'E',-1,64),"float32"
	}
	case "bool":{
		return strconv.FormatBool(v.Bool()),"bool"
	}
	default:{
		log.Fatal("no such type")
	}
	}

	return "",""
}

// 将reflect.Value类型转换成(value.toString() , type)形式
func TransIntoString(v []reflect.Value) ([]string,[]string) {
	log.Println(v[0].Int())
	log.Println(v[0].Type().Name())
	length := len(v)
	Param := make([]string , length)
	Type := make([]string , length)
	for i := 0;i < length ; i++ {
		Param[i], Type[i] = transReflectIntoString(v[i])
	}
	return Param , Type
}