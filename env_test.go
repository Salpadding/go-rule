package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetSetEnv 测试变量 get set
func TestGetSetEnv(t *testing.T) {
	e := newEnv()
	r, _ := parseUnitRule("c0<=10")
	e.set("rule", r)
	assert.Equal(t, r, e.get("rule"))
}

// TestSubEnv 测试子环境
func TestSubEnv(t *testing.T) {
	e := newEnv()
	subEnv := e.newSubEnv()
	r, _ := parseUnitRule("c0<=10")
	e.set("rule", r)
	assert.Equal(t, r, subEnv.get("rule"))
}
