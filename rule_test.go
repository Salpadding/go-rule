package rule

import (
	"testing"
)

func TestUnitRule(t *testing.T) {
	ctx := newContext(map[string]interface{}{
		"c0": 12,
	})
	rs1rule := newUnitRuler(func(c Context) (bool, error) {
		c0, err := c.pluckInt("c0")
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
