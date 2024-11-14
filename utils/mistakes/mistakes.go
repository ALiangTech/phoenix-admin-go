package mistakes

import "fmt"

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
