package rule

import (
	"strconv"
)

const (
	eq  = iota
	neq = iota
	gt  = iota
	gte = iota
	lt  = iota
	lte = iota
)

// Rule 规则
type Rule func(c Context) (bool, error)

// Ruler 拥有规则
type Ruler interface {
	Rule(c Context) (bool, error)
}

func NewUnitRuler(rule Rule) *UnitRuler {
	return &UnitRuler{
		rule: rule,
	}
}

type UnitRuler struct {
	rule Rule
}

func (r *UnitRuler) Rule(c Context) (bool, error) {
	return r.rule(c)
}

func NewComparatorRule(path string, op int, value float64) (*UnitRuler, error) {
	switch op {
	case eq:
		return NewUnitRuler(func(c Context) (bool, error) {
			actual, err := c.PluckFloat64(path)
			if err != nil {
				return false, err
			}
			return actual == value, nil
		}), nil
	case neq:
		return NewUnitRuler(func(c Context) (bool, error) {
			actual, err := c.PluckFloat64(path)
			if err != nil {
				return false, err
			}
			return actual != value, nil
		}), nil
	case gt:
		return NewUnitRuler(func(c Context) (bool, error) {
			actual, err := c.PluckFloat64(path)
			if err != nil {
				return false, err
			}
			return actual > value, nil
		}), nil
	case gte:
		return NewUnitRuler(func(c Context) (bool, error) {
			actual, err := c.PluckFloat64(path)
			if err != nil {
				return false, err
			}
			return actual >= value, nil
		}), nil
	case lt:
		return NewUnitRuler(func(c Context) (bool, error) {
			actual, err := c.PluckFloat64(path)
			if err != nil {
				return false, err
			}
			return actual < value, nil
		}), nil
	case lte:
		return NewUnitRuler(func(c Context) (bool, error) {
			actual, err := c.PluckFloat64(path)
			if err != nil {
				return false, err
			}
			return actual <= value, nil
		}), nil
	}
	return nil, newArgError("arg error: operation " + strconv.Itoa(op) + " is not allowed")
}
