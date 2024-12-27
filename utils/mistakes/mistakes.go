package mistakes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"phoenix-go-admin/routers/model/respond"
)

/*
* @desc 统一报错结构
* 如果不需要继承错误 则传递message 是个空字符串 NewError("", erros.New("xxx"))
 */

func NewError(message string, err error, more ...error) error {
	var format string
	var args []any
	if message != "" {
		format = "%w: %s"
		args = []any{err, message}
	} else {
		format = "%w"
		args = []any{err}
	}

	for _, e := range more {
		format += ": %w"
		args = append(args, e)
	}

	err = fmt.Errorf(format, args...)
	return err
}

func HandleErrorResponse(ctx *gin.Context, code int, err error) {
	msg := StatusText(code)
	if err != nil {
		fmt.Println(NewError(StatusText(code), err).Error())
	}
	ctx.JSON(http.StatusOK, respond.Response{
		Data: nil,
		Code: code,
		Msg:  msg,
	})
}
