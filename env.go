package rule

// env 是环境变量
type env struct {
	parent *env
	vars   map[string]Ruler
}

func newEnv() *env {
	e := new(env)
	e.vars = make(map[string]Ruler)
	return e
}

// find 一层一层往上找变量所在的环境
func (e *env) find(key string) *env {
	if _, ok := e.vars[key]; ok {
		return e
	}
	if e.parent == nil {
		return nil
	}
	return e.parent.find(key)
}

// newSubEnv 新生成一个子环境
func (e *env) newSubEnv() *env {
	ne := newEnv()
	ne.parent = e
	return ne
}

// set 设置变量
func (e *env) set(key string, r Ruler) {
	e.vars[key] = r
}

// get 查找变量
func (e *env) get(key string) Ruler {
	ne := e.find(key)
	if ne == nil {
		return nil
	}
	return ne.vars[key]
}
