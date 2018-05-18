package rule

const (
	typeError = iota
	argError
)

// CustomError 自定义错误
type CustomError struct {
	Code int
	msg  string
}

func (c CustomError) Error() string {
	return c.msg
}

func newTypeError(msg string) CustomError {
	return CustomError{
		Code: typeError,
		msg:  msg,
	}
}

func newArgError(msg string) CustomError {
	return CustomError{
		Code: argError,
		msg:  msg,
	}
}
