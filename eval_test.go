package rule

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 心包经<=-9 或者任意两条经络>=8或<=8
func TestEval(t *testing.T) {
	evaultor := NewEvaluator()
	evaultor.Eval(`(let rs1rule1 'c0>=0')`)
	rs1rule1, err := evaultor.Eval("((rs1rule1))")
	ctx := NewContext(map[string]interface{}{
		"c0": -1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, rs1rule1)
	result, err := rs1rule1.Rule(ctx)
	assert.False(t, result)
}

// 测试变量组合
func TestEvalComposite(t *testing.T) {
	evaultor := NewEvaluator()
	evaultor.Eval(`(let r1 'c0 <= -9')`)
	evaultor.Eval(`(let r2 
			(atLeast 2 'c0>=8' 'c1>=8' 'c2>=8' 'c3>=8' 'c4>=8' 'c5>=8' 'c6>=8' 'c7>=8')
		)`)
	evaultor.Eval(`(let r3 
			(atLeast 2 'c0<=-8' 'c1<=-8' 'c2<=-8' 'c3<=-8' 'c4<=-8' 'c5<=-8' 'c6<=-8' 'c7<=-8')
		)`)
	evaultor.Eval(`(
			let r4
			(or r1 r2 r3)
		)`)
	r1, _ := evaultor.Eval("((r1))")
	r2, _ := evaultor.Eval("r2")
	r3, _ := evaultor.Eval("r3")
	r4, _ := evaultor.Eval("(r4)")
	assert.NotNil(t, r1)
	assert.NotNil(t, r2)
	assert.NotNil(t, r3)
	assert.NotNil(t, r4)
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
	result, err := r4.Rule(ctx)
	assert.NoError(t, err)
	assert.True(t, result)
}

// 测试变量组合
func TestEvalComposite2(t *testing.T) {
	evaultor := NewEvaluator()
	evaultor.Eval(`(let r1 'c0 <= -9')`)
	evaultor.Eval(`(let r1 
			(or r1 (atLeast 2 'c0>=8' 'c1>=8' 'c2>=8' 'c3>=8' 'c4>=8' 'c5>=8' 'c6>=8' 'c7>=8'))
		)`)
	evaultor.Eval(`(let r1
			(or r1 
				(atLeast 2 'c0<=-8' 'c1<=-8' 'c2<=-8' 'c3<=-8' 'c4<=-8' 'c5<=-8' 'c6<=-8' 'c7<=-8')
			)
		)`)
	r1, _ := evaultor.Eval("((r1))")
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
	result, err := r1.Rule(ctx)
	assert.NoError(t, err)
	assert.True(t, result)
}

// 测试 rs1rule1 成功
func Testrs1rule1success(t *testing.T) {
	evaultor := NewEvaluator()
	evaultor.Eval(`(let rs1rule1 'c0 <= -9')`)
	evaultor.Eval(`(let rs1rule1 
			(or rs1rule1 (atLeast 2 'c0>=8' 'c1>=8' 'c2>=8' 'c3>=8' 'c4>=8' 'c5>=8' 'c6>=8' 'c7>=8'))
		)`)
	evaultor.Eval(`(let rs1rule1
			(or rs1rule1 
				(atLeast 2 'c0<=-8' 'c1<=-8' 'c2<=-8' 'c3<=-8' 'c4<=-8' 'c5<=-8' 'c6<=-8' 'c7<=-8')
			)
		)`)
	evaultor.Eval(`(
		let rs1rule1 (and rs1rule1 'isSportOrDrunk=1' )
	)`)
	rs1rule1, _ := evaultor.Eval("((rs1rule1))")
	ctx := NewContext(map[string]interface{}{
		"c0":             10,
		"c1":             0,
		"c2":             0,
		"c3":             0,
		"c4":             0,
		"c5":             0,
		"c6":             0,
		"c7":             0,
		"isSportOrDrunk": 1,
	})
	result, err := rs1rule1.Rule(ctx)
	assert.NoError(t, err)
	assert.True(t, result)
}

// 测试 rs1rule1 失败
func Testrs1rule1fail(t *testing.T) {
	evaultor := NewEvaluator()
	evaultor.Eval(`(let rs1rule1 'c0 <= -9')`)
	evaultor.Eval(`(let rs1rule1 
			(or rs1rule1 (atLeast 2 'c0>=8' 'c1>=8' 'c2>=8' 'c3>=8' 'c4>=8' 'c5>=8' 'c6>=8' 'c7>=8'))
		)`)
	evaultor.Eval(`(let rs1rule1
			(or rs1rule1 
				(atLeast 2 'c0<=-8' 'c1<=-8' 'c2<=-8' 'c3<=-8' 'c4<=-8' 'c5<=-8' 'c6<=-8' 'c7<=-8')
			)
		)`)
	evaultor.Eval(`(
		let rs1rule1 (and rs1rule 'isSportOrDrunk=1')
	`)
	rs1rule1, _ := evaultor.Eval("((rs1rule1))")
	ctx := NewContext(map[string]interface{}{
		"c0":             0,
		"c1":             0,
		"c2":             0,
		"c3":             0,
		"c4":             0,
		"c5":             0,
		"c6":             0,
		"c7":             0,
		"isSportOrDrunk": 0,
	})
	result, err := rs1rule1.Rule(ctx)
	assert.NoError(t, err)
	assert.False(t, result)
}

// 测试 rs2rule1 成功
func Testrs2rule1success(t *testing.T) {
	evaultor := NewEvaluator()
	evaultor.Eval(`(let rs2rule1 (
			and 'c4 >= 20' (
				atLeast 1 'c2>=2' 'c7>=2' 'c3<=-2' 'c6<=-2' 'c5<=-2' 'heartRate >= 80'
			)
		)
	)`)
	evaultor.Eval(`(let rs2rule1 (
			and rs2rule1 'isCold=1'
		)
	)`)
	rs2rule1, _ := evaultor.Eval("((rs1rule1))")
	ctx := NewContext(map[string]interface{}{
		"c0":        0,
		"c1":        0,
		"c2":        0,
		"c3":        0,
		"c4":        20,
		"c5":        0,
		"c6":        0,
		"c7":        0,
		"heartRate": 90,
		"isCold":    1,
	})
	result, err := rs2rule1.Rule(ctx)
	assert.NoError(t, err)
	assert.True(t, result)
}

// 测试 rs2rule1 失败
func Testrs2rule1fail(t *testing.T) {
	evaultor := NewEvaluator()
	evaultor.Eval(`(let rs2rule1 (
			and 'c4 >= 20' (
				atLeast 1 'c2>=2' 'c7>=2' 'c3<=-2' 'c6<=-2' 'c5<=-2' 'heartRate >= 80'
			)
		)
	)`)
	evaultor.Eval(`(let rs2rule1 (
			and rs2rule1 'isCold=1'
		)
	)`)
	rs2rule1, _ := evaultor.Eval("((rs1rule1))")
	ctx := NewContext(map[string]interface{}{
		"c0":        0,
		"c1":        0,
		"c2":        0,
		"c3":        0,
		"c4":        20,
		"c5":        0,
		"c6":        0,
		"c7":        0,
		"heartRate": 0,
		"isCold":    0,
	})
	result, err := rs2rule1.Rule(ctx)
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestFile(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/rules.txt")
	assert.NoError(t, err)
	evaultor := NewEvaluator()
	_, err = evaultor.Eval(string(data))
	assert.NoError(t, err)
	ctx := NewContext(map[string]interface{}{
		"c0":             -8,
		"c1":             0,
		"c2":             0,
		"c3":             0,
		"c4":             0,
		"c5":             0,
		"c6":             0,
		"c7":             0,
		"isSportOrDrunk": 1,
	})
	results := evaultor.Expect(ctx)
	assert.NoError(t, err)
	assert.False(t, results["rs1rule1"].Ok)
}
