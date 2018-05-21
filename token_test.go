package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTokenize 测试 token 生成
func TestTokenize(t *testing.T) {
	program := "(or (and 'c0 <= -9' 'c1 >= -10') 'c3 = 0')"
	tokens := tokenize(program)
	assert.Equal(t, leftParenthesisToken, tokens.shift().tokenType)
	assert.Equal(t, orToken, tokens.shift().tokenType)
	assert.Equal(t, leftParenthesisToken, tokens.shift().tokenType)
	assert.Equal(t, andToken, tokens.shift().tokenType)
	assert.Equal(t, "c0 <= -9", tokens.shift().buf)
}

// TestTokenize 测试 symbol
func TestTokenize2(t *testing.T) {
	program := "var"
	tokens := tokenize(program)
	assert.Equal(t, symbolToken, tokens.shift().tokenType)
}

// TestTokenize 测试 symbol
func TestTokenize3(t *testing.T) {
	program := "(let r1 'c0 <= -9)"
	tokens := tokenize(program)
	assert.Equal(t, leftParenthesisToken, tokens.shift().tokenType)
	assert.Equal(t, letToken, tokens.shift().tokenType)
	assert.Equal(t, symbolToken, tokens.shift().tokenType)
}

