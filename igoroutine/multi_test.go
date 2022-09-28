package igoroutine

import (
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMulti(t *testing.T) {
	multi := NewMulti(3)
	for i := 0; i < 1000; i++ {
		_i := i
		multi.Run(func() error {
			//if _i%100 == 0 {
			//	panic(_i)
			//}
			time.Sleep(2 * time.Millisecond)
			return errors.New(strconv.Itoa(_i))
		})
		//if i%100 == 0 {
		//	t.Log(runtime.NumGoroutine())
		//}
	}
	err := multi.Wait()
	assert.Equal(t, 1000, len(err))
}
