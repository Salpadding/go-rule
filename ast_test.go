package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAst 测试 Ast 生成
func TestAst(t *testing.T) {
	program := "($or ( $and c0<=-9 c1>=-10 ) c3=0)"
	tokens := tokenize(program)
	astNode := buildAST(tokens)
	assert.Equal(t, "", astNode.Token)
	assert.Equal(t, parenthesisNode, astNode.NodeType)
	assert.Equal(t, orNode, astNode.Child.NodeType)
	assert.Equal(t, parenthesisNode, astNode.Child.Next.NodeType)
	assert.Equal(t, leafNode, astNode.Child.Next.Next.NodeType)
}
