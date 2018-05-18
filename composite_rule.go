package rule

// Rule and 方式组合
type AndRuler []Ruler

// Rule 与组合
func (r AndRuler) Rule(c Context) (bool, error) {
	result := true
	for _, ruler := range []Ruler(r) {
		res, err := ruler.Rule(c)
		if err != nil {
			return false, err
		}
		result = result && res
		if !result {
			return result, nil
		}
	}
	return result, nil
}

func NewAndRuler(rules []Ruler) AndRuler {
	return AndRuler(rules)
}

// Rule or 方式组合
type OrRuler []Ruler

// Rule 或组合
func (r OrRuler) Rule(c Context) (bool, error) {
	result := false
	for _, ruler := range []Ruler(r) {
		res, err := ruler.Rule(c)
		if err != nil {
			return false, err
		}
		result = result || res
		if result {
			return result, nil
		}
	}
	return result, nil
}

func NewOrRuler(rules []Ruler) OrRuler {
	return OrRuler(rules)
}

type AtLeastNRuler struct {
	rulers []Ruler
	N      int
}

func (r *AtLeastNRuler) Rule(c Context) (bool, error) {
	n := 0
	for _, rule := range r.rulers {
		res, err := rule.Rule(c)
		if err != nil {
			return false, err
		}
		if res {
			n++
		}
	}
	return n >= r.N, nil
}

func NewAtLeastNRuler(rules []Ruler, n int) *AtLeastNRuler {
	return &AtLeastNRuler{
		rulers: rules,
		N:      n,
	}
}

type AtMostNRuler struct {
	rulers []Ruler
	N      int
}

func (r *AtMostNRuler) Rule(c Context) (bool, error) {
	n := 0
	for _, rule := range r.rulers {
		res, err := rule.Rule(c)
		if err != nil {
			return false, err
		}
		if res {
			n++
		}
	}
	return n <= r.N, nil
}

func NewAtMostNRuler(rules []Ruler, n int) *AtMostNRuler {
	return &AtMostNRuler{
		rulers: rules,
		N:      n,
	}
}
