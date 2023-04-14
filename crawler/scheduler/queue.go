package scheduler

import "learngo/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workChan    chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	// s.workChan = c
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		requestQ := make([]engine.Request, 0)
		workerQ := make([]chan engine.Request, 0)
		for {
			// channel 操作是阻塞的
			// 所有channel操作都放在select里面
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			// send r to a idle worker
			case w := <-s.workChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
