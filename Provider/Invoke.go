package Provider

import (
	"Luka/TaskPool"
	"Luka/util"
	"fmt"
	"log"
	"reflect"
)

var poolInvoke = TaskPool.NewPool(10,10)
var FuncPool map[string]interface{}

func WeakUp(){
	FuncPool = make(map[string]interface{})
	poolInvoke.Start()
}

// 执行函数标号为funcName的函数，参数为v，返回结果转换成[]string
func GetFuncResult(funcName string, v []interface{})([]string, []string){
	nowTask := TaskPool.NewTaskWork(func() interface{} {
		s := FuncPool[funcName]
		funcValue := reflect.ValueOf(s)
		params := util.TransIntoReflect(v)
		log.Println(funcValue)
		res := funcValue.Call(params)
		return res
	})
	poolInvoke.AddTask(nowTask)
	o := nowTask.GetResult()
	// o := nowTask.GetResult()
	return util.TransIntoString(o.([]reflect.Value))
}

func AddFunc(funcName string,tmp interface{}) error {
	if FuncPool[funcName] != nil {
		return fmt.Errorf("there has existed a function")
	}
	//log.Println(tmp)
	FuncPool[funcName] = tmp
	return nil
}