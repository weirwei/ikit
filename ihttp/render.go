package ihttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type DefaultRender struct {
	ErrNo  int         `json:"errNo"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

type Error struct {
	ErrNo  int
	ErrMsg string
}

func (err Error) Error() string {
	return err.ErrMsg
}

func RenderJsonSucc(ctx *gin.Context, data interface{}) {
	var r DefaultRender
	r.ErrNo = 0
	r.ErrMsg = "succ"
	r.Data = data

	ctx.JSON(http.StatusOK, r)
}

func RenderErr(err error) DefaultRender {
	var r DefaultRender

	var code int
	var msg string
	var e Error
	var errorPtr *Error
	switch err = errors.Cause(err); {
	case errors.As(err, &e):
		code = e.ErrNo
		msg = e.ErrMsg
	case errors.As(err, &e):
		code = errorPtr.ErrNo
		msg = errorPtr.ErrMsg
	default:
		code, msg = -1, errors.Cause(err).Error()
	}

	r.ErrNo = code
	r.ErrMsg = msg
	r.Data = gin.H{}
	return r
}

func RenderJsonFail(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, RenderErr(err))
}
