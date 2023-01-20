package semaphore

import (
	"time"
)

type Semaphore struct {
	slots chan struct{}
}

// Creates a Semaphore.
func New(initialCount, maxCount int) *Semaphore {
	sem := &Semaphore{
		slots: make(chan struct{}, maxCount),
	}

	for i := 0; i < initialCount; i++ {
		sem.slots <- struct{}{}
	}

	return sem
}

// Waits for the semaphore to be signaled
// The state of a semaphore object is signaled when its count is greater than zero and
// nonsignaled when its count is equal to zero.
func (sem *Semaphore) Wait() bool {
	<-sem.slots
	return true
}

// Waits for the semaphore to be signaled or timeout
func (sem *Semaphore) TimeWait(d time.Duration) bool {
	select {
	case <-time.After(d):
		return false
	case <-sem.slots:
		return true
	}
}

// Check if the semaphore is signaled
// The state of a semaphore object is signaled when its count is greater than zero and
// nonsignaled when its count is equal to zero
func (sem *Semaphore) TryWait() bool {
	select {
	case <-sem.slots:
		return true
	default:
		return false
	}
}

// Increases the count of the specified semaphore object
func (sem *Semaphore) Post() {
	sem.slots <- struct{}{}
}
