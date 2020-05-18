package TaskPool


type TaskWork struct {
	TaskFunc func() interface{}
	Result   *chan interface{}
}

func NewTaskWork( Func func() interface{}) *TaskWork {
	var ch = make(chan interface{})
	return &TaskWork{
		TaskFunc: Func,
		Result:   &ch,
	}
}

func (Tw *TaskWork) GetResult() interface{}{
	var res interface{}
	select {
	case res = <- *Tw.Result : {
		break
	}
	}
	close(*Tw.Result)
	return res
}

type Pool struct {
	TaskQueue      chan *TaskWork
	RuntimeNumber  int
	Total          int
	isStop         bool
}

// 初始化池子，runtimeNumber是同时允许存在的goroutine数量，
// total是同时允许等待的Task数量
func (p *Pool) Init(runtimeNumber int, total int) {
	p.RuntimeNumber = runtimeNumber
	p.Total = total
	p.isStop = false
	p.TaskQueue = make(chan *TaskWork, total)
}

func (p *Pool) Start() {
	go func(){
		for !p.isStop {
			for i := 0 ; i < p.RuntimeNumber ; i++ {
				select {
				case nowTask,ok := <-p.TaskQueue : {
					if !ok {
						break
					}
					go func() {
						*(nowTask.Result) <- nowTask.TaskFunc()
					}()
					break
				}
				}
			}
		}
	}()
}

// 停止这个池子的使用
func (p *Pool) Stop() {
	p.isStop = true
	close(p.TaskQueue)
}
// 增加一个任务
func (p *Pool) AddTask(task *TaskWork) {
	p.TaskQueue <- task
}

func NewPool(RuntimeNumber int,total int) *Pool {
	ret := &Pool{}
	ret.Init(RuntimeNumber,total)
	return ret
}