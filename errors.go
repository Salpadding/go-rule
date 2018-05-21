package rule

const (
	TypeError = iota
	ArgError
	PropertyNotFoundError
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
		Code: TypeError,
		msg:  msg,
	}
}

func newArgError(msg string) CustomError {
	return CustomError{
		Code: ArgError,
		msg:  msg,
	}
}
