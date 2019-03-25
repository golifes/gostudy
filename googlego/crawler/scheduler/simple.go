package scheduler

import "gostudy/googlego/crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//send request down to worker chan

	go func() {
		s.WorkerChan <- r
	}()
}

func (s *SimpleScheduler) ConfigureWorkChan(c chan engine.Request) {
	s.WorkerChan = c
}
