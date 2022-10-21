package igoroutine

import (
	"sync"
)

// Multi allows start num of goroutines and save err msg.
type Multi struct {
	num     int
	wg      *WaitGroup
	limiter chan struct{}
	lock    sync.Mutex
	errs    []error
}

// exec func with safe mode, prevent crash the server.
func (m *Multi) exec(f func() error) func() error {
	return func() (err error) {
		Safe(func() {
			err = f()
		})()
		return
	}
}

// NewMulti new a Multi with num.
func NewMulti(num int) *Multi {
	if num <= 0 {
		num = 1
	}
	return &Multi{
		num: num,
		wg: &WaitGroup{
			WaitGroup: sync.WaitGroup{},
		},
		limiter: make(chan struct{}, num),
		lock:    sync.Mutex{},
	}
}

// Run func() with limited goroutine num.
func (m *Multi) Run(f func() error) {
	m.limiter <- struct{}{}
	m.wg.Wrap(func() {
		m.appendErr(m.exec(f)())
		<-m.limiter
	})
}

// Wait block until all the missions finished.
// returns err msg.
func (m *Multi) Wait() []error {
	m.wg.Wait()
	return m.errs
}

// appendErr record err during missions with lock.
func (m *Multi) appendErr(err error) {
	if err == nil {
		return
	}
	m.lock.Lock()
	m.errs = append(m.errs, err)
	m.lock.Unlock()
}
