package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetSetEnv 测试变量 get set
func TestGetSetEnv(t *testing.T) {
	e := newEnv()
	r, _ := parseUnitRule("c0<=10")
	e.setLocal("rule", r)
	v, err := e.get("rule")
	assert.NoError(t, err)
	assert.Equal(t, r, v)
}
