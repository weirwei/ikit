package ihttp

import (
	"fmt"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/weirwei/ikit/igoroutine"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

const (
	SSENameMsg      = "msg"
	SSENameError    = "error"
	SSENameClose    = "close"
	SSENameResponse = "response"
)

type sendEvents struct {
	ctx *gin.Context

	writer chan any      // 数据通道，使用管道保证线程安全
	over   chan struct{} // 用于阻塞流程，stream 发完之后才能 return
	config *sseConfig
}

type sseConfig struct {
	eventId uint64
	timeout time.Duration
}

type SSEOption func(c *sseConfig)

func NewSendEvents(ctx *gin.Context, opts ...SSEOption) *sendEvents {
	s := &sendEvents{
		ctx:    ctx,
		writer: make(chan any),
		over:   make(chan struct{}),
		config: &sseConfig{
			eventId: 0,
			timeout: 2 * time.Second,
		},
	}
	for _, opt := range opts {
		opt(s.config)
	}
	go igoroutine.Safe(func() {
		defer close(s.over) // 数据推完，通知结束，解除阻塞
		for {
			select {
			case data, ok := <-s.writer:
				if !ok { // 通道关系，直接结束
					return
				}
				s.pushData(data)
			}
		}
	})()
	return s
}

func SetTimeout(timeout time.Duration) SSEOption {
	return func(c *sseConfig) {
		c.timeout = timeout
	}
}

func (s *sendEvents) Send(data any) error {
	timer := time.NewTimer(s.config.timeout)
	select {
	case s.writer <- data:
	case <-timer.C: // 推送超时
		return fmt.Errorf("send events timeout")
	}
	timer.Stop()
	return nil
}

func (s *sendEvents) End() {
	close(s.writer)
	<-s.over // 阻塞流程
}

func (s *sendEvents) pushData(data any) {
	var event sse.Event
	switch v := data.(type) {
	case sse.Event:
		event = v
		id, _ := strconv.ParseUint(event.Id, 0, 64)
		if id < s.config.eventId { // 需要保证消息id递增
			event.Id = strconv.FormatUint(s.config.eventId, 10)
		}
	case error:
		event = sse.Event{
			Id:    strconv.FormatUint(s.config.eventId, 10),
			Event: SSENameError,
			Data:  v,
		}
	default:
		event = sse.Event{
			Id:    strconv.FormatUint(s.config.eventId, 10),
			Event: SSENameMsg,
			Data:  v,
		}
	}
	s.ctx.Render(http.StatusOK, event)
	atomic.AddUint64(&s.config.eventId, 1) // 发完消息后，id+1
}

func (s *sendEvents) Close() {
	s.ctx.Render(http.StatusOK, sse.Event{Event: SSENameClose})
}
