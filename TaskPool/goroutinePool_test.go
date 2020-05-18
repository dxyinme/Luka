package TaskPool

import (
	"log"
	"testing"
)


var poolTest = &Pool{
	isStop: true,
}

func add(a,b int) int {
	return a + b
}

func TestPool_AddTask(t *testing.T) {

}

func TestPool_Init(t *testing.T) {
	poolTest.Init(10,10)
}

func TestPool_SetFinishCallback(t *testing.T) {

}

func TestPool_Start(t *testing.T) {
	poolTest.Init(10,10)
	poolTest.Start()
	for i := 0 ; i < 10 ; i++ {
		log.Println(i)
		now := NewTaskWork(func() interface{} {
			return add(i,i)
		})
		poolTest.AddTask(now)
		log.Printf("result %d : %d \n",i,now.GetResult().(int))
	}
	poolTest.Stop()
}

func TestPool_Stop(t *testing.T) {

}

func TestTaskWork_GetResult(t *testing.T) {

}