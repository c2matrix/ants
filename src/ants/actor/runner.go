package actor

//运行器(基本不用管了)
type IRunner interface {
	Close()
	Send(interface{}) bool
	LoopActor(IActor)
}

//简单的测试
type actorRunner struct {
	buffChannel
}

func newRunner() IRunner {
	this := new(actorRunner)
	this.init(1000)
	return this
}

//interfaces
func (this *actorRunner) LoopActor(ator IActor) {
	this.Loop(func(data interface{}) {
		ator.OnMessage(data)
	})
}

func (this *actorRunner) Send(data interface{}) bool {
	return this.Push(data)
}