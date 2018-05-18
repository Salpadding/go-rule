package rule

import (
	"strings"
)

const (
	typeError = iota
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

// given a map, pull a property from it at some deeply nested depth
// this reimplements (most of) JS `pluck` in go: https://github.com/gjohnson/pluck
func pluck(o map[string]interface{}, path string) interface{} {
	// support dots for now ebcause thats all we need
	parts := strings.Split(path, ".")

	if len(parts) == 1 && o[parts[0]] != nil {
		// if there is only one part, just return that property value
		return o[parts[0]]
	} else if len(parts) > 1 && o[parts[0]] != nil {
		var prev map[string]interface{}
		var ok bool
		if prev, ok = o[parts[0]].(map[string]interface{}); !ok {
			// not an object type! ...or a map, yeah, that.
			return nil
		}

		for i := 1; i < len(parts)-1; i += 1 {
			// we need to check the existence of another
			// map[string]interface for every property along the way
			cp := parts[i]

			if prev[cp] == nil {
				// didn't find the property, it's missing
				return nil
			}
			var ok bool
			if prev, ok = prev[cp].(map[string]interface{}); !ok {
				return nil
			}
		}

		if prev[parts[len(parts)-1]] != nil {
			return prev[parts[len(parts)-1]]
		} else {
			return nil
		}
	}

	return nil
}

// Context 规则引擎的输入上下文
type Context struct {
	data       map[string]interface{}
	intPool    map[string]int
	stringPool map[string]string
	boolPool   map[string]bool
}

func newContext(data map[string]interface{}) Context {
	return Context{
		data:       data,
		intPool:    make(map[string]int),
		stringPool: make(map[string]string),
		boolPool:   make(map[string]bool),
	}
}

func (c Context) pluck(path string) interface{} {
	return pluck(c.data, path)
}

// 解析出整数
func (c Context) pluckInt(path string) (int, error) {
	i, ok := c.intPool[path]
	if ok {
		return i, nil
	}
	v := c.pluck(path)
	i, ok = v.(int)
	if ok {
		c.intPool[path] = i
		return i, nil
	}
	return 0, newTypeError("type error: " + path + " is not integer")
}

// 解析出整数数组
func (c Context) pluckInts(paths []string) ([]int, error) {
	res := make([]int, len(paths))
	for idx, path := range paths {
		i, err := c.pluckInt(path)
		if err != nil {
			return nil, err
		}
		res[idx] = i
	}
	return res, nil
}

// 设值
func (c Context) setBool(path string, val bool) {
	c.boolPool[path] = val
}

func (c Context) pluckBool(path string) (bool, error) {
	b, ok := c.boolPool[path]
	if ok {
		return b, nil
	}
	v := c.pluck(path)
	b, ok = v.(bool)
	if ok {
		c.boolPool[path] = b
		return b, nil
	}
	return false, newTypeError("type error: " + path + " is not bool")
}
