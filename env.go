package rule

import "errors"

// env 是环境变量
type env struct {
	parent  *env
	locals  map[string]Ruler
	globals map[string]Ruler
}

func newEnv() *env {
	e := new(env)
	e.locals = make(map[string]Ruler)
	e.globals = make(map[string]Ruler)
	return e
}

// newSubEnv 新生成一个子环境
func (e *env) newSubEnv() *env {
	ne := newEnv()
	ne.parent = e
	return ne
}

// setLocal 设置变量
func (e *env) setLocal(key string, r Ruler) {
	e.locals[key] = r
}

// setGlobal 设置变量
func (e *env) setGlobal(key string, r Ruler) {
	e.globals[key] = r
}

// get 查找变量
func (e *env) get(key string) (Ruler, error) {
	r, ok := e.locals[key]
	if ok {
		return r, nil
	}
	r, ok = e.globals[key]
	if ok {
		return r, nil
	}
	return nil, errors.New("cannot find variable" + key)
}
