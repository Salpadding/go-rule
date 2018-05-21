package rule

import(
	"testing"
	"github.com/stretchr/testify/assert"	
)

// TestAst 测试 Ast 生成
func TestAst(t *testing.T) {
	program := "(or ( and `c0 <= -9` `c1 >= -10`) `c3 = 0`)"
	tokens := tokenize(program)
	astNode := buildAST(tokens)
	assert.Equal(t, orToken, astNode.child.token.tokenType)
}

// TestAst2 测试 Ast 生成
func TestAst2(t *testing.T) {
	program := "(let var (or ( and `c0 <= -9` `c1 >= -10`) `c3 = 0`))"
	tokens := tokenize(program)
	astNode := buildAST(tokens)
	assert.Equal(t, letToken, astNode.child.token.tokenType)
	assert.Equal(t, symbolToken, astNode.child.next.token.tokenType)	
}