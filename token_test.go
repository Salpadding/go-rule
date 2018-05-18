package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTokenize 测试 token 生成
func TestTokenize(t *testing.T) {
	program := "($or ( $and c0<=-9 c1>=-10 ) c3=0 )"
	tokens := tokenize(program)
	assert.Equal(t, "(", tokens.shift())
	assert.Equal(t, "$or", tokens.shift())
	assert.Equal(t, "(", tokens.shift())
	assert.Equal(t, "$and", tokens.shift())
	assert.Equal(t, "c0<=-9", tokens.shift())
}
