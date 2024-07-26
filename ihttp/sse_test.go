package ihttp

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weirwei/ikit/igoroutine"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendEvents_Send(t *testing.T) {
	eng := gin.New()
	t.Run("test", func(t *testing.T) {
		eng.Handle(http.MethodGet, "", func(ctx *gin.Context) {
			event := NewSendEvents(ctx)
			defer func() {
				event.End()
				event.Close()
			}()
			event.Send("hello")
			event.Send("hello1")
			event.Send("hello2")
			muti := igoroutine.NewMulti(3)
			for i := 0; i < 10; i++ {
				muti.Run(func() error {
					_i := i
					event.Send(fmt.Sprintf("go hello %d", _i))
					return nil
				})
			}
			muti.Wait()
		})
		srv := httptest.NewServer(eng)
		response, err := srv.Client().Get(srv.URL)
		if err != nil {
			t.Error(err)
			return
		}
		defer response.Body.Close()
		reader := bufio.NewReader(response.Body)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				if errors.Is(err, io.EOF) {
					t.Log("end")
					break
				}
			}
			t.Log(string(line))
		}
	})
}
