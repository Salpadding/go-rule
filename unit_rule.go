package rule

// Rule 规则
type Rule func(c Context) (bool, error)

// Ruler 拥有规则
type Ruler interface {
	Rule(c Context) (bool, error)
}

func newUnitRuler(rule Rule) *UnitRuler {
	return &UnitRuler{
		Rule: rule,
	}
}

type UnitRuler struct {
	Rule Rule
}
