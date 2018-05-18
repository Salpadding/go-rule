package rule

import (
	"testing"
)

func TestUnitRule(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 12,
	})
	rs1rule := NewUnitRuler(func(c Context) (bool, error) {
		c0, err := c.PluckInt("c0")
		if err != nil {
			return false, err
		}
		return c0 == 12, nil
	})
	res, err := rs1rule.Rule(Context(ctx))
	if err != nil || !res {
		panic("fail")
	}
}

func TestComparatorGTERule(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 12,
	})
	rule, _ := NewComparatorRule("c0", gte, 10)
	res, err := rule.Rule(ctx)
	if err != nil || !res {
		panic(err)
	}
}

func TestComparatorNERule(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 12,
	})
	rule, _ := NewComparatorRule("c0", neq, 10)
	res, err := rule.Rule(ctx)
	if err != nil || !res {
		panic(err)
	}
}

func TestComparatorEQRule(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 12,
	})
	rule, _ := NewComparatorRule("c0", eq, 12)
	res, err := rule.Rule(ctx)
	if err != nil || !res {
		panic(err)
	}
}

func TestComparatorGTRule(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 12,
	})
	rule, _ := NewComparatorRule("c0", gt, 1)
	res, err := rule.Rule(ctx)
	if err != nil || !res {
		panic(err)
	}
}

func TestComparatorLTRule(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 12,
	})
	rule, _ := NewComparatorRule("c0", lt, 1000)
	res, err := rule.Rule(ctx)
	if err != nil || !res {
		panic(err)
	}
}

func TestComparatorLTERule(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 12,
	})
	rule, _ := NewComparatorRule("c0", lte, 12)
	res, err := rule.Rule(ctx)
	if err != nil || !res {
		panic(err)
	}
}
