package igoroutine

import (
	"sync"
	"testing"
)

func TestSafe(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		Safe(func() {
			panic(1)
		})()
	}()
}
