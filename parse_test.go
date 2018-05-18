package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAst 测试 Ast 生成
func TestParseUnitRule(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 12,
	})
	ruler, err := parseUnitRule("c0=12")
	assert.NoError(t, err)
	res, err := ruler.Rule(ctx)
	assert.NoError(t, err)
	assert.True(t, res)
}

// TestParseRule 测试解析规则
func TestParseRule1(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": -9,
		"c1": 0,
		"c2": 0,
		"c3": 0,
		"c4": 0,
		"c5": 0,
		"c6": 0,
		"c7": 0,
	})
	ruler, err := ParseRule(`($or c0<=-9  
			($atLeast 2 c0>=80  c1>=8 c2>=8 c3>=8 c4>=8 c5>=8 c6>=8 c7>=8) 
			($atLeast 2 c0<=-8  c1<=-8 c2<=-8 c3<=-8 c4<=-8 c5<=-8 c6<=-8 c7<=-8)   			  
		)`)
	assert.NoError(t, err)
	res, err := ruler.Rule(ctx)
	assert.NoError(t, err)
	assert.True(t, res)
}

func TestParseRule2(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 0,
		"c1": 0,
		"c2": 0,
		"c3": 0,
		"c4": 0,
		"c5": 0,
		"c6": 0,
		"c7": 0,
	})
	ruler, err := ParseRule(`($or c0<=-9  
			($atLeast 2 c0>=80  c1>=8 c2>=8 c3>=8 c4>=8 c5>=8 c6>=8 c7>=8) 
			($atLeast 2 c0<=-8  c1<=-8 c2<=-8 c3<=-8 c4<=-8 c5<=-8 c6<=-8 c7<=-8)   			  
		)`)
	assert.NoError(t, err)
	res, err := ruler.Rule(ctx)
	assert.NoError(t, err)
	assert.False(t, res)
}

func TestParseRule3(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 0,
		"c1": 8,
		"c2": 0,
		"c3": 0,
		"c4": 0,
		"c5": 8,
		"c6": 0,
		"c7": 0,
	})
	ruler, err := ParseRule(`($or c0<=-9  
			($atLeast 2 c0>=80  c1>=8 c2>=8 c3>=8 c4>=8 c5>=8 c6>=8 c7>=8) 
			($atLeast 2 c0<=-8  c1<=-8 c2<=-8 c3<=-8 c4<=-8 c5<=-8 c6<=-8 c7<=-8)   			  
		)`)
	assert.NoError(t, err)
	res, err := ruler.Rule(ctx)
	assert.NoError(t, err)
	assert.True(t, res)
}

func TestParseRule4(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 0,
		"c1": -9,
		"c2": 0,
		"c3": 0,
		"c4": 0,
		"c5": -9,
		"c6": 0,
		"c7": 0,
	})
	ruler, err := ParseRule(`($or c0<=-9  
			($atLeast 2 c0>=80  c1>=8 c2>=8 c3>=8 c4>=8 c5>=8 c6>=8 c7>=8) 
			($atLeast 2 c0<=-8  c1<=-8 c2<=-8 c3<=-8 c4<=-8 c5<=-8 c6<=-8 c7<=-8)   			  
		)`)
	assert.NoError(t, err)
	res, err := ruler.Rule(ctx)
	assert.NoError(t, err)
	assert.True(t, res)
}
