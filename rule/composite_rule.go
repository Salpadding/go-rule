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

func newAndRuler(rules []Ruler) AndRuler {
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

func newOrRuler(rules []Ruler) OrRuler {
	return OrRuler(rules)
}
