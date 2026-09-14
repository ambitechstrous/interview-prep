package atomic

import "sync"

type AtomicInteger struct {
	val int
	mu  *sync.Mutex
}

type IAtomicInteger interface {
	Get() int
	Add(num int)
}

func NewAtomicInteger(val int) IAtomicInteger {
	return &AtomicInteger{
		val: val,
		mu:  &sync.Mutex{},
	}
}

func (a *AtomicInteger) Get() int {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.val
}

func (a *AtomicInteger) Add(num int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.val += num
}
