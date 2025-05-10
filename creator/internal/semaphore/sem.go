package semaphore

type Semaphore struct {
	sem chan struct{}
}

func NewSemaphore(maxConc int) Semaphore {
	return Semaphore{
		sem: make(chan struct{}, maxConc),
	}
}

func (s Semaphore) Acquire() {
	s.sem <- struct{}{}
}

func (s Semaphore) Release() {
	<-s.sem
}
