package scheduler

import "learngo/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// 会卡死 用 go func
	go func() {
		// 直接分发 request 我们无法控制速率 =>  使用队列
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

// func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
// 	s.workerChan = c
// }
