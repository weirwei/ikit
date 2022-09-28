package igoroutine

import (
	"log"
	"runtime/debug"
	"sync"
)

// WaitGroup contains sync.WaitGroup
type WaitGroup struct {
	sync.WaitGroup
}

// Wrap start a goroutine with safe mode
func (wg *WaitGroup) Wrap(f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		Safe(func() {
			f()
		})()
	}()
}

// Safe mode, recover the panic, prevent crash the server.
func Safe(f func()) func() {
	return func() {
		defer func() {
			if rec := recover(); rec != nil {
				stack := debug.Stack()
				log.Printf("panic: %v, stack: %s", rec, stack)
			}
		}()
		f()
	}
}
