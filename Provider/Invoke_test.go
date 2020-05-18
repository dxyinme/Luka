package Provider

import (
	"Luka/util"
	"testing"
)

func TestAddFunc(t *testing.T) {

}

func TestGetFuncResult(t *testing.T) {
	WeakUp()
	_ = AddFunc("Add", func(a,b int) int { return a + b })
	Param := []string{"1","15"}
	Type := []string{"int","int"}
	interParam,_ := util.TransformList(Param,Type)
	_,_ = GetFuncResult("Add",interParam)
}

func TestWeakUp(t *testing.T) {

}

func Test_transIntoReflect(t *testing.T) {

}